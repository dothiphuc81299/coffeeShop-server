package categoryimpl

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	store      *store
	drinkStore drink.Store
}

func NewService(store *store, drinkStore drink.Store) category.Service {
	return &service{
		store:      store,
		drinkStore: drinkStore,
	}
}

func (s *service) Create(ctx context.Context, cmd *category.CategoryBody) (err error) {
	if s.checkNameExisted(ctx, cmd.Name) {
		return category.ErrCategoryExisted
	}

	entity := category.CategoryRaw{
		ID:           primitive.NewObjectID(),
		Name:         cmd.Name,
		SearchString: format.NonAccentVietnamese(cmd.Name),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	err = s.store.InsertOne(ctx, entity)
	if err != nil {
		return err
	}

	return err
}

func (s *service) checkNameExisted(ctx context.Context, name string) bool {
	total := s.store.CountByCondition(ctx, bson.M{"name": name})
	return total > 0
}

func (s *service) ListAll(ctx context.Context, q *category.SearchCategoryQuery) ([]category.CategoryRaw, int64, error) {
	var (
		wg      sync.WaitGroup
		mutex   sync.Mutex
		cond    = bson.M{}
		total   int64
		res     = make([]category.CategoryRaw, 0)
		errList []error
	)

	if q.Name != "" {
		cond["searchString"] = format.NonAccentVietnamese(q.Name)
	}

	if q.Limit == 0 {
		q.Limit = 10
	}
	if q.Page == 0 {
		q.Page = 1
	}

	skip := (q.Page - 1) * q.Limit

	findOpts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(q.Limit))

	wg.Add(2)

	go func() {
		defer wg.Done()
		count := s.store.CountByCondition(ctx, cond)
		mutex.Lock()
		total = count
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		categories, err := s.store.FindByCondition(ctx, cond, findOpts)
		if err != nil {
			mutex.Lock()
			errList = append(errList, err)
			mutex.Unlock()
			return
		}
		mutex.Lock()
		res = categories
		mutex.Unlock()
	}()

	wg.Wait()

	if len(errList) > 0 {
		return nil, 0, fmt.Errorf("error(s) occurred: %v", errList)
	}

	return res, total, nil
}

func (s *service) GetDetail(ctx context.Context, id primitive.ObjectID) (category.CategoryRaw, error) {
	return s.store.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (s *service) DeleteCategory(ctx context.Context, id primitive.ObjectID) error {
	cate, err := s.store.FindOneByCondition(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if cate.ID.IsZero() {
		return category.ErrCategoryNotFound
	}

	if err := s.store.DeleteByID(ctx, cate.ID); err != nil {
		return err
	}

	// delete menu by category id
	if err := s.drinkStore.DeleteByCategoryID(ctx, cate.ID); err != nil {
		return err
	}
	return nil
}

func (s *service) Update(ctx context.Context, id primitive.ObjectID, body category.CategoryBody) (err error) {
	cate, err := s.store.FindOneByCondition(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if cate.ID.IsZero() {
		return category.ErrCategoryNotFound
	}

	cate.Name = body.Name
	cate.SearchString = format.NonAccentVietnamese(body.Name)
	cate.UpdatedAt = time.Now().UTC()

	err = s.store.UpdateByID(ctx, cate.ID, bson.M{"$set": cate})
	if err != nil {
		return err
	}

	return err
}
