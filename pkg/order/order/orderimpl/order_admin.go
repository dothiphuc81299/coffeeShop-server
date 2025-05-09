package orderimpl

// const success = "success"

// type OrderAdminService struct {
// 	OrderDAO model.OrderDAO
// 	DrinkDAO model.DrinkDAO
// 	UserDAO  model.UserDAO
// }

// func NewOrderAdminService(d *model.CommonDAO) model.OrderAdminService {
// 	return &OrderAdminService{
// 		OrderDAO: d.Order,
// 		DrinkDAO: d.Drink,
// 		UserDAO:  d.User,
// 	}
// }

// func (o *OrderAdminService) SearchByStatus(ctx context.Context, query model.CommonQuery) ([]model.OrderResponse, int64) {
// 	var (
// 		cond     = bson.M{}
// 		total    int64
// 		wg       sync.WaitGroup
// 		res      = make([]model.OrderResponse, 0)
// 		condUser = bson.M{}
// 	)

// 	// assign
// 	query.AssignStatus(&cond)
// 	query.AssignUsername(&condUser)

// 	total = o.OrderDAO.CountByCondition(ctx, cond)
// 	orders, _ := o.OrderDAO.FindByCondition(ctx, cond, query.GetFindOptsUsingPageOne())

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

// func (o *OrderAdminService) ChangeStatus(ctx context.Context, order model.OrderRaw, status model.StatusBody, staff model.Staff) (res string, err error) {
// 	if order.Status != "pending" {
// 		return "", errors.New(locale.OrderStatusIsInvalid)
// 	}
// 	payload := bson.M{
// 		"updatedAt": time.Now(),
// 		"status":    status.Status,
// 		"updatedBy": staff.ID,
// 	}

// 	err = o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
// 	if err != nil {
// 		return res, err
// 	}

// 	return status.Status, err
// }

// func (o *OrderAdminService) UpdateOrderSuccess(ctx context.Context, order model.OrderRaw, staff model.Staff) error {
// 	if order.Status != "pending" {
// 		return errors.New(locale.OrderStatusCanNotUpdate)
// 	}
// 	payload := bson.M{
// 		"updatedAt": time.Now(),
// 		"status":    "success",
// 		"updatedBy": staff.ID,
// 	}

// 	err := o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
// 	if err != nil {
// 		return err
// 	}

// 	//get user
// 	user, err := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})
// 	if err != nil {
// 		return err
// 	}
// 	var currentPointUpdate float64
// 	if !order.IsPoint {
// 		// calculate currentPoint
// 		if (order.TotalPrice >= 30000) && (order.TotalPrice) <= 50000 {
// 			currentPointUpdate = user.CurrentPoint + 1
// 		} else if (order.TotalPrice > 50000) && (order.TotalPrice) <= 100000 {
// 			currentPointUpdate = user.CurrentPoint + 2
// 		} else if order.TotalPrice > 100000 {
// 			currentPointUpdate = user.CurrentPoint + 3
// 		}
// 	}
// 	if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 		return errors.New(locale.UpdatePointFailed)
// 	}

// 	return nil
// }

// func (o *OrderAdminService) CancelOrder(ctx context.Context, order model.OrderRaw, staff model.Staff) error {
// 	if order.Status != "pending" {
// 		return errors.New(locale.OrderStatusCanNotUpdate)
// 	}
// 	payload := bson.M{
// 		"updatedAt": time.Now(),
// 		"status":    "cancel",
// 		"updatedBy": staff.ID,
// 	}

// 	err := o.OrderDAO.UpdateByID(ctx, order.ID, bson.M{"$set": payload})
// 	if err != nil {
// 		return err
// 	}
// 	//get user
// 	user, err := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})
// 	if err != nil {
// 		return err
// 	}

// 	if order.IsPoint && order.Point > 0 {
// 		currentPointUpdate := user.CurrentPoint + order.Point
// 		if err = o.UserDAO.UpdateByID(ctx, user.ID, bson.M{"$set": bson.M{"currentPoint": currentPointUpdate}}); err != nil {
// 			return errors.New(locale.UpdatePointFailed)
// 		}
// 	}

// 	return nil
// }

// func (o *OrderAdminService) FindByID(ctx context.Context, id primitive.ObjectID) (model.OrderRaw, error) {
// 	return o.OrderDAO.FindOneByCondition(ctx, bson.M{"_id": id})
// }

// func (o *OrderAdminService) GetDetail(ctx context.Context, order model.OrderRaw) (doc model.OrderResponse) {
// 	user, _ := o.UserDAO.FindOneByCondition(ctx, bson.M{"_id": order.User})
// 	if user.ID.IsZero() {
// 		return
// 	}
// 	userInfo := user.GetUserInfo()

// 	res := order.GetResponse(userInfo, order.Drink, order.Status)
// 	return res
// }

// func (o *OrderAdminService) GetStatistic(ctx context.Context, query model.CommonQuery) (model.StatisticResponse, error) {
// 	var (
// 		cond = bson.M{
// 			"status": "success",
// 		}
// 		res = model.StatisticResponse{}
// 	)

// 	query.AssignStartAtAndEndAtByStatistic(&cond)

// 	orders, err := o.OrderDAO.FindByCondition(ctx, cond, query.GetFindOptionsUsingSort())

// 	if err != nil {
// 		return model.StatisticResponse{}, err
// 	}
// 	var tempResutl = make([]model.StatisticByDrink, 0)

// 	for _, i := range orders {
// 		for _, drink := range i.Drink {
// 			if !o.checkDuplicate(drink, tempResutl) {
// 				dr := model.StatisticByDrink{
// 					ID:            drink.ID,
// 					Name:          drink.Name,
// 					TotalQuantity: float64(drink.Quantity),
// 					TotalSale:     float64(drink.Quantity) * drink.Price,
// 				}
// 				tempResutl = append(tempResutl, dr)
// 			}
// 		}
// 	}

// 	sort.Slice(tempResutl, func(i, j int) bool {
// 		return tempResutl[j].TotalQuantity < tempResutl[i].TotalQuantity
// 	})

// 	var result = make([]model.StatisticByDrink, 0)
// 	result = tempResutl
// 	if len(tempResutl) > 4 {
// 		result = tempResutl[:4]
// 	}

// 	for _, i := range tempResutl {
// 		res.TotalQuantity += i.TotalQuantity
// 		res.TotalSale += i.TotalSale
// 	}

// 	res.Statistic = result

// 	return res, nil
// }

// func (s *OrderAdminService) checkDuplicate(drink model.DrinkInfo, tempResutl []model.StatisticByDrink) bool {
// 	for k, i := range tempResutl {
// 		if i.ID == drink.ID {
// 			tempResutl[k].TotalQuantity = tempResutl[k].TotalQuantity + float64(drink.Quantity)
// 			tempResutl[k].TotalSale = tempResutl[k].TotalSale + float64(drink.Quantity)*drink.Price
// 			return true
// 		}
// 	}
// 	return false
// }
