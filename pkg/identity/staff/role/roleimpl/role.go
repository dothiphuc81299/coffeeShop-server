package roleimpl

import (
	"context"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type service struct {
	store      *store
	staffStore staff.Store
}

func NewService(store *store, staffStore staff.Store) role.Service {
	return &service{
		store:      store,
		staffStore: staffStore,
	}
}

func (s *service) FindByID(ctx context.Context, id primitive.ObjectID) (role.StaffRoleRaw, error) {
	return s.store.FindByID(ctx, id)
}

func (s *service) Delete(ctx context.Context, id primitive.ObjectID) error {
	data, err := s.store.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.store.DeleteByID(ctx, data.ID)
}

func (s *service) Update(ctx context.Context, cmd role.UpdateStaffRoleCommand) error {
	payload := bson.M{
		"$set": bson.M{
			"name":        cmd.Name,
			"permissions": cmd.Permissions,
			"updatedAt":   time.Now().UTC(),
		},
	}

	data, err := s.store.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = s.store.UpdateByID(ctx, data.ID, payload)
	if err != nil {
		return role.ErrCanNotUpdate
	}

	data.Name = cmd.Name
	data.Permissions = cmd.Permissions

	cond := bson.M{
		"role": data.ID,
	}

	payloadStaff := bson.M{
		"$set": bson.M{
			"permissions": cmd.Permissions,
			"updatedAt":   time.Now().UTC(),
		},
	}
	err = s.staffStore.UpdateBycondition(ctx, cond, payloadStaff)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListStaffRole(ctx context.Context, q *query.CommonQuery) ([]role.StaffRoleRaw, int64) {
	if q.Limit == 0 {
		q.Limit = 10
	}

	if q.Page == 0 {
		q.Page = 1
	}

	skip := (q.Page - 1) * q.Limit
	cond := bson.M{}

	docs, err := s.store.FindByCondition(ctx, cond, options.Find().SetLimit(q.Limit).SetSkip(skip))
	if err != nil {
		return nil, 0
	}

	total := s.store.CountByCondition(ctx, cond)
	return docs, total
}

func (s *service) Create(ctx context.Context, body role.CreateStaffRoleCommand) error {
	now := time.Now().UTC()
	roleNewBSON := role.StaffRoleRaw{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: body.Permissions,
	}

	err := s.store.InsertOne(ctx, &roleNewBSON)
	if err != nil {
		return err
	}
	return nil
}
