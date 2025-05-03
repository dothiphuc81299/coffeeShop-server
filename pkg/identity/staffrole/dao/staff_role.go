package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const staffRoleCol = "staff-roles"

// StaffRoleDAO ...
type StaffRoleDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewStaffRoleDAO ...
func NewStaffRoleDAO(db *mongo.Database) model.StaffRoleDAO {
	return &StaffRoleDAO{
		DB:  db,
		Col: db.Collection(staffRoleCol),
	}
}

// CountByCondition ...
func (ud *StaffRoleDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := ud.Col.CountDocuments(ctx, cond)
	return total
}

// FindOneByCondition ...
func (ud *StaffRoleDAO) FindOneByCondition(ctx context.Context, cond interface{}) (u model.StaffRoleRaw, err error) {
	err = ud.Col.FindOne(ctx, cond).Decode(&u)
	return u, err
}

// FindByID ...
func (ud *StaffRoleDAO) FindByID(ctx context.Context, id model.AppID) (model.StaffRoleRaw, error) {
	cond := bson.M{"_id": id}
	return ud.FindOneByCondition(ctx, cond)
}

// FindByCondition ...
func (ud *StaffRoleDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.StaffRoleRaw, err error) {
	cursor, err := ud.Col.Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

// InsertOne ...
func (ud *StaffRoleDAO) InsertOne(ctx context.Context, u *model.StaffRoleRaw) error {
	_, err := ud.Col.InsertOne(ctx, u)
	return err
}

// UpdateByID ...
func (ud *StaffRoleDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := ud.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

func (w *StaffRoleDAO) DeleteByID(ctx context.Context, id model.AppID) error {
	_, err := w.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
