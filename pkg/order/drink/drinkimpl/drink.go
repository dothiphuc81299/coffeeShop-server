package drinkimpl

import (
	"context"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	store         *store
	categoryStore category.Store
}

func NewService(store *store, categoryStore category.Store) drink.Service {
	return &service{
		store:         store,
		categoryStore: categoryStore,
	}
}

func (s *service) Create(ctx context.Context, body drink.DrinkBody) (err error) {
	if s.checkNameExisted(ctx, body.Name) {
		return drink.ErrDrinkNameExisted
	}

	payload := body.NewDrinkRaw()
	err = s.store.InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return err
}

func (s *service) checkNameExisted(ctx context.Context, name string) bool {
	total := s.store.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

func (s *service) ListAll(ctx context.Context, q query.CommonQuery) ([]drink.DrinkAdminResponse, int64) {
	var (
		wg    sync.WaitGroup
		cond  = bson.M{}
		res   = make([]drink.DrinkAdminResponse, 0)
		total int64
	)

	q.AssignKeyword(&cond)
	q.AssignActive(&cond)
	q.AssignCategory(&cond)

	total = s.store.CountByCondition(ctx, cond)
	drinks, _ := s.store.FindByCondition(ctx, cond, q.GetFindOptsUsingPageOne())
	if len(drinks) > 0 {
		wg.Add(len(drinks))
		res = make([]drink.DrinkAdminResponse, len(drinks))
		for index, value := range drinks {
			go func(abc drink.DrinkRaw, i int) {
				defer wg.Done()
				cat, _ := s.categoryStore.FindOneByCondition(ctx, bson.M{"_id": abc.Category})
				catTemp := drink.CategoryGetInfo(cat)
				temp := abc.DrinkGetAdminResponse(catTemp)
				res[i] = temp
			}(value, index)

		}

		wg.Wait()
	}

	return res, total
}

func (s *service) Update(ctx context.Context, id primitive.ObjectID, body drink.DrinkBody) (err error) {
	data, err := s.store.FindOneByCondition(ctx, id)
	if err != nil {
		return err
	}

	doc := body.NewDrinkRaw()

	// assign
	data.Name = doc.Name
	data.SearchString = doc.SearchString
	data.Category = doc.Category
	data.Price = doc.Price
	data.Image = doc.Image

	err = s.store.UpdateByID(ctx, data.ID, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) FindByID(ctx context.Context, id primitive.ObjectID) (drink.DrinkRaw, error) {
	return s.store.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (s *service) ChangeStatus(ctx context.Context, drinkID primitive.ObjectID) (status bool, err error) {
	result, err := s.store.FindOneByCondition(ctx, bson.M{"_id": drinkID})
	if err != nil {
		return false, err
	}

	if result.ID.IsZero() {
		return false, drink.ErrDrinkNotFound
	}

	active := !result.Active
	payload := bson.M{
		"$set": bson.M{
			"active": active,
		},
	}
	err = s.store.UpdateByID(ctx, result.ID, payload)
	if err != nil {
		return
	}
	return active, nil
}

func (s *service) GetDetail(ctx context.Context, data drink.DrinkRaw) drink.DrinkAdminResponse {
	cat, _ := s.categoryStore.FindOneByCondition(ctx, bson.M{"_id": data.Category})
	catTemp := drink.CategoryGetInfo(cat)
	temp := data.DrinkGetAdminResponse(catTemp)
	return temp
}

func (s *service) DeleteDrink(ctx context.Context, drinkID primitive.ObjectID) error {
	return s.store.DeleteByID(ctx, drinkID)
}
