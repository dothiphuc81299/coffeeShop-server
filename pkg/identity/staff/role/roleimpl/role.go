package roleimpl

import (
	"context"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (s *service) Update(ctx context.Context, id primitive.ObjectID, body role.CreateStaffRoleCommand) error {
	payload := bson.M{
		"$set": bson.M{
			"name":        body.Name,
			"permissions": body.Permissions,
			"updatedAt":   time.Now().UTC(),
		},
	}

	data, err := s.store.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.store.UpdateByID(ctx, data.ID, payload)
	if err != nil {
		return role.ErrCanNotUpdate
	}

	data.Name = body.Name
	data.Permissions = body.Permissions

	cond := bson.M{
		"role": data.ID,
	}

	payloadStaff := bson.M{
		"$set": bson.M{
			"permissions": body.Permissions,
			"updatedAt":   time.Now().UTC(),
		},
	}
	err = s.staffStore.UpdateBycondition(ctx, cond, payloadStaff)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ListStaffRole(ctx context.Context, q query.CommonQuery) ([]role.StaffRoleRaw, int64) {
	cond := bson.M{}
	docs, err := s.store.FindByCondition(ctx, cond)
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
