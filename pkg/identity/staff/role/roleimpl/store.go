package roleimpl

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/infra/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const staffRoleCol = "staff-roles"

type store struct {
	Col *mongo.Collection
}

func Newstore(db mongodb.DBConnector) *store {
	return &store{
		Col: db.GetCollection(staffRoleCol),
	}
}

func (ud *store) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := ud.Col.CountDocuments(ctx, cond)
	return total
}

func (ud *store) FindOneByCondition(ctx context.Context, cond interface{}) (u role.StaffRoleRaw, err error) {
	err = ud.Col.FindOne(ctx, cond).Decode(&u)
	return u, err
}

func (ud *store) FindByID(ctx context.Context, id primitive.ObjectID) (role.StaffRoleRaw, error) {
	cond := bson.M{"_id": id}
	return ud.FindOneByCondition(ctx, cond)
}

func (ud *store) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []role.StaffRoleRaw, err error) {
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

func (ud *store) InsertOne(ctx context.Context, u *role.StaffRoleRaw) error {
	_, err := ud.Col.InsertOne(ctx, u)
	return err
}

func (ud *store) UpdateByID(ctx context.Context, id primitive.ObjectID, payload interface{}) error {
	_, err := ud.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

func (w *store) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := w.Col.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
