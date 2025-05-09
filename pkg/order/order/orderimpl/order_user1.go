package orderimpl

// type OrderAppService struct {
// 	OrderDAO         model.OrderDAO
// 	DrinkDAO         model.DrinkDAO
// 	UserDAO          model.UserDAO
// 	DrinkAnalyticDAO model.DrinkAnalyticDAO
// }

// const POINT = 3000

// func NewOrderAppService(d *model.CommonDAO) model.OrderAppService {
// 	return &OrderAppService{
// 		OrderDAO:         d.Order,
// 		DrinkDAO:         d.Drink,
// 		UserDAO:          d.User,
// 		DrinkAnalyticDAO: d.DrinkAnalytic,
// 	}
// }

// func (s *service) Create(ctx context.Context, cmd order.OrderBody) (doc order.OrderResponse, err error) {
// 	// convert order payload
// 	drinks := make([]order.DrinkInfo, 0)
// 	for _, value := range cmd.Drink {
// 		drinkID, _ := primitive.ObjectIDFromHex(value.Name)
// 		drinkRaw, err := s.drinkStore.FindOneByCondition(ctx, bson.M{"_id": drinkID})
// 		if err != nil {
// 			return doc, err
// 		}
// 		drink := order.DrinkInfo{
// 			ID:       drinkID,
// 			Name:     drinkRaw.Name,
// 			Price:    drinkRaw.Price,
// 			Quantity: value.Quantity,
// 		}
// 		doc.TotalPrice += drink.Price * float64(drink.Quantity)
// 		drinks = append(drinks, drink)
// 	}
// 	var currentPointUpdate float64
// 	// if !order.IsPoint {
// 	// 	// calculate currentPoint
// 	// 	if (doc.TotalPrice >= 30000) && (doc.TotalPrice) <= 50000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 1
// 	// 	} else if (doc.TotalPrice > 50000) && (doc.TotalPrice) <= 100000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 2
// 	// 	} else if doc.TotalPrice > 100000 {
// 	// 		currentPointUpdate = user.CurrentPoint + 3
// 	// 	}
// 	// }

// 	s.calculate()

// 	orderPayload := order.NewOrderRaw(user.ID, drinks, doc.TotalPrice)

// 	err = s.store.InsertOne(ctx, orderPayload)
// 	if err != nil {
// 		return doc, errors.New(locale.OrderKeyCanNotCreateOrder)
// 	}

// 	if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 		return doc, errors.New(locale.UpdatePointFailed)
// 	}

// 	userInfo := user.GetUserInfo()
// 	res := orderPayload.GetResponse(userInfo, drinks, orderPayload.Status)
// 	return res, nil

// }

func (s *service) calculate() {
	// if cmd.Point > 0 {
	// 	if user.CurrentPoint < cmd.Point {
	// 		return doc, errors.New(locale.CurrentPointIsNotEnough)
	// 	}
	// 	missPoint := cmd.Point * POINT

	// 	if missPoint > doc.TotalPrice {
	// 		point := missPoint - doc.TotalPrice
	// 		point = point / POINT
	// 		///	currentPointUpdate = point
	// 		currentPointUpdate = user.CurrentPoint - cmd.Point + point
	// 		cmd.Point = cmd.Point - point
	// 		doc.TotalPrice = 0
	// 	} else {
	// 		doc.TotalPrice -= missPoint
	// 		currentPointUpdate = user.CurrentPoint - cmd.Point
	// 	}
	// }
}

// func (s *service) FindByID(ctx context.Context, id primitive.ObjectID) (model.OrderRaw, error) {
// 	return s.store.FindOneByCondition(ctx, bson.M{"_id": id})
// }

// func (s *service) GetDetail(ctx context.Context, order model.OrderRaw) (doc model.OrderResponse) {
// 	user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})

// 	if user.ID.IsZero() {
// 		return
// 	}

// 	userInfo := user.GetUserInfo()
// 	res := order.GetResponse(userInfo, order.Drink, order.Status)
// 	return res
// }

// func (o *OrderAppService) Search(ctx context.Context, query model.CommonQuery, user model.UserRaw) ([]model.OrderResponse, int64) {
// 	var (
// 		cond = bson.M{
// 			"user": user.ID,
// 		}
// 		total int64
// 		wg    sync.WaitGroup
// 		res   = make([]model.OrderResponse, 0)
// 	)
// 	// assign
// 	query.AssignStatus(&cond)
// 	total = o.OrderDAO.CountByCondition(ctx, cond)
// 	orders, _ := o.OrderDAO.FindByCondition(ctx, cond, query.GetFindOptsUsingPage())

// 	if len(orders) > 0 {
// 		wg.Add(len(orders))
// 		res = make([]model.OrderResponse, len(orders))
// 		for index, order := range orders {
// 			go func(od model.OrderRaw, i int) {
// 				defer wg.Done()
// 				user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": od.User})

// 				userInfo := user.GetUserInfo()

// 				temp := od.GetResponse(userInfo, od.Drink, od.Status)
// 				res[i] = temp

// 			}(order, index)
// 		}
// 		wg.Wait()

// 	}
// 	return res, total
// }

// func (o *OrderAppService) RejectOrder(ctx context.Context, user model.UserRaw, order model.OrderRaw) error {
// 	now := time.Now()
// 	createdAt := order.CreatedAt
// 	then := createdAt.Add(time.Duration(+2) * time.Minute)

// 	if now.After(then) || order.Status != "pending" {
// 		return errors.New(locale.OrderCanNotCancel)
// 	}

// 	payload := bson.M{
// 		"updatedAt": time.Now(),
// 		"status":    "cancel",
// 		"updatedBy": user.ID,
// 	}

// 	err := o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
// 	if err != nil {
// 		return errors.New(locale.CommonKeyErrorWhenHandle)
// 	}

// 	// check lai diem
// 	currentPointUpdate := user.CurrentPoint + order.Point
// 	if order.IsPoint && order.Point > 0 {
// 		if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 			return errors.New(locale.UpdatePointFailed)
// 		}
// 	}

// 	return nil

// }
