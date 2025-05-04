package dao

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const questionCol = "questions"

// QuestionDAO ....
type QuestionDAO struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewQuestionDAO ...
func NewQuestionDAO(db *mongo.Database) model.QuestionDAO {
	return &QuestionDAO{
		DB:  db,
		Col: db.Collection(questionCol),
	}
}

// InsertOne ...
func (d *QuestionDAO) InsertOne(ctx context.Context, u model.QuestionRaw) error {
	_, err := d.Col.InsertOne(ctx, u)
	return err
}

// FindByCondition ...
func (w *QuestionDAO) FindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.QuestionRaw, err error) {
	cursor, err := w.Col.Find(ctx, cond, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}
	return docs, nil
}

// CountByCondition ...
func (w *QuestionDAO) CountByCondition(ctx context.Context, cond interface{}) int64 {
	total, _ := w.Col.CountDocuments(ctx, cond)
	return total
}

// UpdateByID ...
func (w *QuestionDAO) UpdateByID(ctx context.Context, id model.primitive.ObjectID, payload interface{}) error {
	_, err := w.Col.UpdateOne(ctx, bson.M{"_id": id}, payload)
	return err
}

// FindOneByCondition ...
func (w *QuestionDAO) FindOneByCondition(ctx context.Context, cond interface{}) (doc model.QuestionRaw, err error) {
	err = w.Col.FindOne(ctx, cond).Decode(&doc)
	return
}
