package dao

// import (
// 	"context"
// 	"fmt"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const ordercol = "orders"

// // OrderDAO ....
// type OrderDAO struct {
// 	DB  *mongo.Database
// 	Col *mongo.Collection
// }

// // NewOrderDAO ...
// func NewOrderDAO(db *mongo.Database) model.OrderDAO {
// 	return &OrderDAO{
// 		DB:  db,
// 		Col: db.Collection(ordercol),
// 	}
// }

// // InsertOne ...
// func (d *OrderDAO) InsertOne(ctx context.Context, u model.OrderRaw) error {
// 	_, err := d.Col.InsertOne(ctx, u)
// 	return err
// }

// // FindByCondition ...
// func (d *OrderDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.OrderRaw, err error) {
// 	cursor, err := d.Col.Find(ctx, cond, opts...)

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer cursor.Close(ctx)
// 	if err := cursor.All(ctx, &docs); err != nil {

// 		return nil, err
// 	}

// 	return docs, nil
// }

// // CountByCondition ...
// func (w *OrderDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
// 	total, _ := w.Col.CountDocuments(ctx, cond)
// 	return total
// }

// // UpdateByID ...
// func (w *OrderDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
// 	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
// 	return err
// }

// // FindOneByCondition ...
// func (w *OrderDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.OrderRaw, err error) {
// 	err = w.Col.FindOne(ctx, cond).Decode(&doc)
// 	return
// }

// // AggregateOrder
// func (w *OrderDAO) AggregateOrder(ctx context.Context, cond interface{}) ([]*model.StatisticByDrink, error) {
// 	var (
// 		results = make([]*model.StatisticByDrink, 0)
// 	)
// 	match := bson.M{
// 		"$match": cond,
// 	}

// 	project := bson.M{
// 		"$project": bson.M{
// 			"_id":           1,
// 			"totalQuantity": 1,
// 		},
// 	}

// 	group := bson.M{
// 		"$group": bson.M{
// 			"_id": "$drink._id",
// 			"totalQuantity": bson.M{
// 				"$sum": "$drink.quantity",
// 			},
// 		},
// 	}

// 	cursor, err := w.Col.Aggregate(ctx, []bson.M{match, project, group})
// 	fmt.Println("cur", cursor)
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 		return results, nil
// 	}

// 	defer cursor.Close(ctx)
// 	err = cursor.All(ctx, &results)
// 	fmt.Println("results", results)
// 	return results, err
// }
