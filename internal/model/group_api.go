package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuizGroupCommon struct {
	ID            primitive.ObjectID   `json:"_id"`
	Name          string               `json:"name"`
	Quizzes       []primitive.ObjectID `json:"quizzes"`
	Active        bool                 `json:"active"`
	CreatedAt     TimeResponse         `json:"createdAt"`
	UpdatedAt     TimeResponse         `json:"updatedAt"`
	TotalQuestion float64              `bson:"totalQuestion"`
}
type QuizGroupBody struct {
	Name    string   `json:"name"`
	Quizzes []string `json:"quizzes"`
	Active  bool     `json:"active"`
}

func (l QuizGroupBody) Validate() (err error) {
	err = validation.ValidateStruct(&l,
		validation.Field(&l.Name, validation.Required.Error("GroupNameIsRequired")),
		validation.Field(&l.Quizzes, validation.Required.Error("QuizIDIsRequired"),
			validation.Each(is.MongoID.Error("QuizIDInValid"))),
	)
	if err != nil {
		return err
	}

	return nil
}

// QuizGroupNewBSON ...
func (l QuizGroupBody) QuizGroupNewBSON() QuizGroupRaw {
	now := time.Now()
	quizzes := make([]primitive.ObjectID, 0)
	for _, value := range l.Quizzes {
		id, _ := primitive.ObjectIDFromHex(value)
		quizzes = append(quizzes, id)
	}
	totalQuestion := len(quizzes)
	return QuizGroupRaw{
		ID:            primitive.NewObjectID(),
		Name:          l.Name,
		Quizzes:       quizzes,
		Active:        l.Active,
		TotalQuestion: float64(totalQuestion),
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// QuizGroupGetCommon ...
func QuizGroupGetCommon(l QuizGroupRaw) QuizGroupCommon {
	return QuizGroupCommon{
		ID:            l.ID,
		Active:        l.Active,
		Quizzes:       l.Quizzes,
		Name:          l.Name,
		TotalQuestion: l.TotalQuestion,
		UpdatedAt:     TimeResponse{Time: l.UpdatedAt},
		CreatedAt:     TimeResponse{Time: l.CreatedAt},
	}
}
