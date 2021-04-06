package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const staffCol = "staffs"

// StaffDAO ...
type StaffDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewStaffDAO ...
func NewStaffDAO(db *mongo.Database) model.StaffDAO {
	return &StaffDAO{
		DB:  db,
		Col: db.Collection(staffCol),
	}
}

// CountByCondition ...
func (ud *StaffDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := ud.Col.CountDocuments(ctx, cond)
	return total
}

// FindOneByCondition ...
func (ud *StaffDAO) FindOneByCondition(ctx context.Context, cond interface{}) (u model.StaffRaw, err error) {
	err = ud.Col.FindOne(ctx, cond).Decode(&u)
	return u, err
}

// FindByID ...
func (ud *StaffDAO) FindByID(ctx context.Context, id model.AppID) (model.StaffRaw, error) {
	cond := bson.M{"_id": id}
	return ud.FindOneByCondition(ctx, cond)
}

// FindByCondition ...
func (ud *StaffDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.StaffRaw, err error) {
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
func (ud *StaffDAO) InsertOne(ctx context.Context, u model.StaffRaw) error {
	_, err := ud.Col.InsertOne(ctx, u)
	return err
}

// UpdateByID ...
func (ud *StaffDAO) UpdateByID(ctx context.Context, id model.AppID, payload interface{}) error {
	_, err := ud.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}
