package dao

// import (
// 	"context"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const staffCol = "staffs"

// // StaffDAO ...
// type StaffDAO struct {
// 	DB  *mongo.Database
// 	Col *mongo.Collection
// }

// // NewStaffDAO ...
// func NewStaffDAO(db *mongo.Database) model.StaffDAO {
// 	return &StaffDAO{
// 		DB:  db,
// 		Col: db.Collection(staffCol),
// 	}
// }

// // CountByCondition ...
// func (ud *StaffDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
// 	total, _ := ud.Col.CountDocuments(ctx, cond)
// 	return total
// }

// // FindOneByCondition ...
// func (ud *StaffDAO) FindOneByCondition(ctx context.Context, cond interface{}) (u model.Staff, err error) {
// 	err = ud.Col.FindOne(ctx, cond).Decode(&u)
// 	return u, err
// }

// // FindByID ...
// func (ud *StaffDAO) FindByID(ctx context.Context, id model.primitive.ObjectID) (model.Staff, error) {
// 	cond := bson.M{"_id": id}
// 	return ud.FindOneByCondition(ctx, cond)
// }

// // FindByCondition ...
// func (ud *StaffDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.Staff, err error) {
// 	cursor, err := ud.Col.Find(ctx, cond, opts...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)
// 	if err := cursor.All(ctx, &docs); err != nil {
// 		return nil, err
// 	}
// 	return docs, nil
// }

// // InsertOne ...
// func (ud *StaffDAO) InsertOne(ctx context.Context, u model.Staff) error {
// 	_, err := ud.Col.InsertOne(ctx, u)
// 	return err
// }

// // UpdateByID ...
// func (ud *StaffDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
// 	_, err := ud.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
// 	return err
// }

// func (ud *StaffDAO) UpdateBycondition(ctx context.Context, cond interface{}, payload interface{}) error {
// 	_, err := ud.Col.UpdateOne(ctx, cond, payload)
// 	return err
// }

// func (w *StaffDAO) DeleteByID(ctx context.Context, id model.primitive.ObjectID) error {
// 	_, err := w.Col.DeleteOne(ctx, bson.M{"_id": id})
// 	return err
// }
