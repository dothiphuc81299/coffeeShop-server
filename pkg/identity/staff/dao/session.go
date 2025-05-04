package dao

// import (
// 	"context"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const sessionCol = "sessions"

// // SessionDAO ...
// type SessionDAO struct {
// 	DB  *mongo.Database
// 	Col *mongo.Collection
// }

// // NewSessionDAO ...
// func NewSessionDAO(db *mongo.Database) model.SessionDAO {
// 	return &SessionDAO{
// 		DB:  db,
// 		Col: db.Collection(sessionCol),
// 	}
// }

// // CountByCondition ...
// func (ud *SessionDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
// 	total, _ := ud.Col.CountDocuments(ctx, cond)
// 	return total
// }

// // FindOneByCondition ...
// func (ud *SessionDAO) FindOneByCondition(ctx context.Context, cond interface{}) (u model.SessionRaw, err error) {
// 	err = ud.Col.FindOne(ctx, cond).Decode(&u)
// 	return u, err
// }

// // FindByID ...
// func (ud *SessionDAO) FindByID(ctx context.Context, id model.primitive.ObjectID) (model.SessionRaw, error) {
// 	cond := bson.M{"_id": id}
// 	return ud.FindOneByCondition(ctx, cond)
// }

// // FindByCondition ...
// func (ud *SessionDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.SessionRaw, err error) {
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
// func (ud *SessionDAO) InsertOne(ctx context.Context, u model.SessionRaw) error {
// 	_, err := ud.Col.InsertOne(ctx, u)
// 	return err
// }

// // UpdateByID ...
// func (ud *SessionDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
// 	_, err := ud.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
// 	return err
// }

// // RemoveByCondition ...
// func (ud *SessionDAO) RemoveByCondition(ctx context.Context, cond interface{}) error {
// 	_, err := ud.Col.DeleteMany(ctx, cond)
// 	return err
// }
