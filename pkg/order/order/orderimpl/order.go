package orderimpl

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/order"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/shippingaddress"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/useraccount"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	store         *store
	drinkStore    drink.Store
	userStore     useraccount.Store
	shippingStore shippingaddress.Store
}

func NewService(store *store, drinkStore drink.Store, userStore useraccount.Store, shippingStore shippingaddress.Store) order.Service {
	return &service{
		store:         store,
		drinkStore:    drinkStore,
		userStore:     userStore,
		shippingStore: shippingStore,
	}
}

func (s *service) Create(ctx context.Context, cmd order.OrderBody) (doc order.OrderResponse, err error) {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return doc, errors.New("account is invalid")
	}

	userAcc, err := s.userStore.FindOneByCondition(ctx, bson.M{"user_id": account.ID})
	if err != nil {
		return doc, err
	}

	// 1.prepare list drink
	orderItems := make([]*order.OrderItemRaw, 0)
	for _, value := range cmd.Items {
		drinkRaw, err := s.drinkStore.FindOneByCondition(ctx, bson.M{"_id": value.DrinkID})
		if err != nil {
			return doc, err
		}

		item := order.OrderItemRaw{
			DrinkID:  value.DrinkID,
			Name:     drinkRaw.Name,
			Price:    drinkRaw.Price,
			Quantity: value.Quantity,
			Total:    float64(value.Quantity) * drinkRaw.Price,
		}
		doc.TotalPrice += item.Total
		orderItems = append(orderItems, &item)
	}

	// 2. calculate currentPoint (if user use point)
	var currentPointUpdate float64 = userAcc.CurrentPoint
	if !cmd.UsePoint {
		switch {
		case doc.TotalPrice >= 30000 && doc.TotalPrice <= 50000:
			currentPointUpdate += 1
		case doc.TotalPrice > 50000 && doc.TotalPrice <= 100000:
			currentPointUpdate += 2
		case doc.TotalPrice > 100000:
			currentPointUpdate += 3
		}
	}

	// 3. hanle shipping
	var address string
	var shippingAddressID primitive.ObjectID
	if cmd.ShippingAddressID != "" {
		shippingAddressID, err = primitive.ObjectIDFromHex(cmd.ShippingAddressID)
		if err != nil {
			return doc, errors.New("invalid shipping address ID")
		}

		shippingAddress, err := s.shippingStore.FindOneByCondition(ctx, bson.M{"_id": shippingAddressID})
		if err != nil {
			return doc, errors.New("unable to find shipping address")
		}
		address = shippingAddress.Address
	} else {
		newAddress := shippingaddress.UserShippingAddressRaw{
			ID:        primitive.NewObjectID(),
			UserID:    userAcc.ID,
			FullName:  cmd.Shipping.FullName,
			Phone:     cmd.Shipping.Phone,
			Address:   cmd.Shipping.Address,
			Province:  cmd.Shipping.Province,
			City:      cmd.Shipping.City,
			Ward:      cmd.Shipping.Ward,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
		err = s.shippingStore.InsertOne(ctx, newAddress)
		if err != nil {
			return doc, errors.New("unable to save shipping address")
		}
		shippingAddressID = newAddress.ID
		address = newAddress.Address
	}

	now := time.Now().UTC()
	orderRaw := order.OrderRaw{
		ID:                primitive.NewObjectID(),
		UserID:            userAcc.ID,
		ShippingAddressID: shippingAddressID,
		Items:             orderItems,
		Status:            "pending",
		Total:             doc.TotalPrice,
		CreatedAt:         now,
		UpdatedAt:         now,
		UsePoint:          cmd.UsePoint,
		Point:             cmd.Point,
	}

	err = s.store.InsertOne(ctx, orderRaw)
	if err != nil {
		return doc, order.ErrCannotCreateOrder
	}

	// 5.Update point
	err = s.userStore.UpdateOne(ctx, userAcc.ID, bson.M{"$set": bson.M{"current_point": currentPointUpdate}})
	if err != nil {
		return doc, order.ErrUpdatePointFailed
	}

	res := orderRaw.GetResponse(order.UserInfo{
		ID:       userAcc.ID,
		UserName: userAcc.LoginName,
		Address:  address,
	}, orderItems, orderRaw.Status, nil)

	return res, nil
}

func (s *service) GetDetail(ctx context.Context, id primitive.ObjectID) (doc order.OrderResponse) {
	orderRaw, err := s.store.FindOneByCondition(ctx, bson.M{"_id": id})
	if err != nil || orderRaw.ID.IsZero() {
		return
	}

	user, _ := s.userStore.FindOneByCondition(ctx, bson.M{"_id": orderRaw.UserID})
	if user.ID.IsZero() {
		return
	}

	shippingAddr, _ := s.shippingStore.FindOneByCondition(ctx, bson.M{"_id": orderRaw.ShippingAddressID})

	doc = orderRaw.GetResponse(order.UserInfo{
		ID:       user.ID,
		UserName: user.LoginName,
	}, orderRaw.Items, orderRaw.Status, &shippingAddr)
	return
}

func (s *service) Search(ctx context.Context, query *order.SearchOrdersQuery) ([]order.OrderResponse, int64) {
	var (
		cond  = bson.M{}
		total int64
		wg    sync.WaitGroup
		res   = make([]order.OrderResponse, 0)
	)

	opts := options.Find()
	if query.Limit > 0 {
		opts.SetLimit(query.Limit).SetSkip((query.Page) * query.Limit)
	}

	if query.UserID != "" {
		cond["user_id"] = query.UserID
	}

	if query.OrderID != "" {
		cond["_id"] = query.OrderID
	}

	if query.Status != "" {
		cond["status"] = query.Status
	}

	total = s.store.CountByCondition(ctx, cond)
	orders, _ := s.store.FindByCondition(ctx, cond, opts)

	if len(orders) > 0 {
		wg.Add(len(orders))
		res = make([]order.OrderResponse, len(orders))

		for index, value := range orders {
			go func(od order.OrderRaw, i int) {
				defer wg.Done()

				userAcc, _ := s.userStore.FindOneByCondition(ctx, bson.M{"_id": od.UserID})
				if userAcc.ID.IsZero() {
					return
				}

				shippingAddr, _ := s.shippingStore.FindOneByCondition(ctx, bson.M{"_id": od.ShippingAddressID})

				res[i] = od.GetResponse(order.UserInfo{
					ID:       userAcc.ID,
					UserName: userAcc.LoginName,
				}, od.Items, od.Status, &shippingAddr)
			}(value, index)
		}
		wg.Wait()
	}
	return res, total
}

func (s *service) RejectOrder(ctx context.Context, cmd *order.UpdateOrderStatusCommand) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.User {
		return errors.New("account is invalid")
	}

	orderRaw, err := s.store.FindOneByCondition(ctx, bson.M{"_id": cmd.ID})
	if err != nil {
		return order.ErrOrderNotFound
	}

	if orderRaw.Status != "pending" {
		return order.ErrOrderCanNotCancel
	}

	if time.Since(orderRaw.CreatedAt) > 2*time.Minute {
		return order.ErrOrderCanNotCancel
	}

	update := bson.M{
		"status":     "cancel",
		"updated_at": time.Now().UTC(),
	}

	if err := s.store.UpdateByID(ctx, cmd.ID, bson.M{"$set": update}); err != nil {
		return err
	}

	if orderRaw.UsePoint && orderRaw.Point > 0 {
		user, err := s.userStore.FindOneByCondition(ctx, bson.M{"_id": orderRaw.UserID})
		if err != nil {
			return order.ErrUserNotFound
		}

		newPoint := user.CurrentPoint + orderRaw.Point
		if err := s.userStore.UpdateOne(ctx, user.ID, bson.M{"$set": bson.M{"current_point": newPoint}}); err != nil {
			return order.ErrUpdatePointFailed
		}
	}

	return nil
}

func (s *service) UpdateOrderSuccess(ctx context.Context, cmd *order.UpdateOrderStatusCommand) error {
	if cmd.Status != "success" {
		return order.ErrOrderStatusCanNotUpdate
	}

	result, err := s.store.FindOneByCondition(ctx, bson.M{"_id": cmd.ID})
	if err != nil {
		return err
	}
	if result.Status != "pending" {
		return order.ErrOrderStatusCanNotUpdate
	}

	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.Staff {
		return errors.New("unauthorized")
	}

	update := bson.M{
		"status":     "success",
		"updated_at": time.Now().UTC(),
	}

	if err := s.store.UpdateByID(ctx, result.ID, bson.M{"$set": update}); err != nil {
		return err
	}

	user, err := s.userStore.FindOneByCondition(ctx, bson.M{"_id": result.UserID})
	if err != nil {
		return err
	}

	if !result.UsePoint {
		newPoint := user.CurrentPoint
		switch {
		case result.Total >= 30000 && result.Total <= 50000:
			newPoint += 1
		case result.Total > 50000 && result.Total <= 100000:
			newPoint += 2
		case result.Total > 100000:
			newPoint += 3
		}
		if err := s.userStore.UpdateOne(ctx, user.ID, bson.M{"$set": bson.M{"current_point": newPoint}}); err != nil {
			return order.ErrUpdatePointFailed
		}
	}

	return nil
}

func (s *service) GetStatistic(ctx context.Context, cmd query.CommonQuery) (order.StatisticResponse, error) {
	var (
		cond = bson.M{
			"status": "success",
		}
		res     = order.StatisticResponse{}
		tempRes = make([]order.StatisticByDrink, 0)
	)

	cmd.AssignStartAtAndEndAtByStatistic(&cond)

	orders, err := s.store.FindByCondition(ctx, cond, cmd.GetFindOptionsUsingSort())
	if err != nil {
		return order.StatisticResponse{}, err
	}

	for _, ord := range orders {
		for _, item := range ord.Items {
			if item == nil {
				continue
			}
			s.aggregateDrinkStatistic(item, &tempRes)
		}
	}

	sort.Slice(tempRes, func(i, j int) bool {
		return tempRes[j].TotalQuantity < tempRes[i].TotalQuantity
	})

	// Lấy top 4 món bán chạy
	top := tempRes
	if len(tempRes) > 4 {
		top = tempRes[:4]
	}

	// Tính tổng số lượng & doanh thu
	for _, i := range tempRes {
		res.TotalQuantity += i.TotalQuantity
		res.TotalSale += i.TotalSale
	}
	res.Statistic = top

	return res, nil
}

// Hợp nhất số liệu nếu trùng món
func (s *service) aggregateDrinkStatistic(item *order.OrderItemRaw, list *[]order.StatisticByDrink) {
	for i := range *list {
		if (*list)[i].ID == item.DrinkID {
			(*list)[i].TotalQuantity += float64(item.Quantity)
			(*list)[i].TotalSale += float64(item.Quantity) * item.Price
			return
		}
	}
	*list = append(*list, order.StatisticByDrink{
		ID:            item.DrinkID,
		Name:          item.Name,
		TotalQuantity: float64(item.Quantity),
		TotalSale:     float64(item.Quantity) * item.Price,
	})
}
