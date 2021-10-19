package model

import (
	"sort"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	QuestionBody struct {
		Order    int                   `json:"order"`
		Question string                `json:"question"`
		Answers  []QuestionAnswersBody `json:"answers"`
	}

	// QuestionAnswersBody ...
	QuestionAnswersBody struct {
		Answer  string `json:"answer"`
		Correct bool   `json:"correct"`
		Order   int    `json:"order"`
	}

	QuestionCommon struct {
		ID        primitive.ObjectID      `json:"_id"`
		Order     int                     `json:"order"`
		Question  string                  `json:"question"`
		Answers   []QuestionAnswersCommon `json:"answers"`
		Active    bool                    `json:"active"`
		CreatedAt TimeResponse            `json:"createdAt"`
		UpdatedAt TimeResponse            `json:"updatedAt"`
	}

	// QuestionAnswersCommon  ...
	QuestionAnswersCommon struct {
		ID      primitive.ObjectID `json:"_id" `
		Answer  string             `json:"answer"`
		Correct bool               `json:"correct"`
		Order   int                `json:"order"`
	}

	// QuestionBodyUpdate ...
	QuestionBodyUpdate struct {
		Order    int                         `json:"order"`
		Question string                      `json:"question"`
		Answers  []QuestionAnswersBodyUpdate `json:"answers"`
	}

	// QuestionAnswersBodyUpdate ...
	QuestionAnswersBodyUpdate struct {
		ID      string `json:"_id"`
		Answer  string `json:"answer"`
		Correct bool   `json:"correct"`
		Order   int    `json:"order"`
	}
)

// Validate ...
func (l QuestionBody) Validate() error {
	err := validation.ValidateStruct(&l,
		validation.Field(&l.Order, validation.Required.Error("Order is Required")),
		validation.Field(&l.Question, validation.Required.Error("QuestionIsRequired")),
		validation.Field(&l.Answers, validation.Length(2, 5).Error("AnswersLengthMustBeGreaterThanOneAndSmallerThanFive")),
	)

	if err != nil {
		return err
	}

	err = validation.Validate(l.Answers)
	return err
}

func (l QuestionBodyUpdate) Validate() error {
	err := validation.ValidateStruct(&l,
		validation.Field(&l.Order, validation.Required.Error("OrderIsRequired")),
		validation.Field(&l.Question, validation.Required.Error("QuestionIsRequired")),
		validation.Field(&l.Answers, validation.Length(2, 5).Error("AnswersLengthMustBeGreaterThanOneAndSmallerThanFive")),
	)

	if err != nil {
		return err
	}

	err = validation.Validate(l.Answers)
	return err
}

// Validate ...
func (l QuestionAnswersBody) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Answer, validation.Required.Error("AnswerIsRequired")),
		validation.Field(&l.Order, validation.Required.Error("AnswerOrderIsRequired")),
	)
}

// QuestionNewBSON ...
func (l QuestionBody) QuestionNewBSON() QuestionRaw {
	now := time.Now()
	answers := make([]QuestionAnswersBSON, 0)
	for _, i := range l.Answers {
		answer := i.AnswersNewBSON()
		answers = append(answers, answer)
	}

	return QuestionRaw{
		ID:        primitive.NewObjectID(),
		Order:     l.Order,
		Question:  l.Question,
		Answers:   answers,
		Active:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// QuestionAnswersBody ...
func (l QuestionAnswersBody) AnswersNewBSON() QuestionAnswersBSON {
	return QuestionAnswersBSON{
		ID:      primitive.NewObjectID(),
		Answer:  l.Answer,
		Correct: l.Correct,
		Order:   l.Order,
	}
}

// QuestionGetCommonAPI ...
func QuestionGetCommonAPI(p QuestionRaw) QuestionCommon {
	answers := make([]QuestionAnswersCommon, 0)

	sort.Slice(p.Answers, func(i, j int) bool {
		return p.Answers[i].Order < p.Answers[j].Order
	})

	for _, i := range p.Answers {

		answer := QuestionAnswerGetCommonAPI(i)
		answers = append(answers, answer)
	}

	result := QuestionCommon{
		ID:        p.ID,
		Question:  p.Question,
		Order:     p.Order,
		Answers:   answers,
		UpdatedAt: TimeResponse{Time: p.UpdatedAt},
		CreatedAt: TimeResponse{Time: p.CreatedAt},
		Active:    p.Active,
	}
	return result
}

// QuestionAnswerGetCommonAPI ...
func QuestionAnswerGetCommonAPI(p QuestionAnswersBSON) QuestionAnswersCommon {
	return QuestionAnswersCommon{
		ID:      p.ID,
		Answer:  p.Answer,
		Correct: p.Correct,
		Order:   p.Order,
	}
}

// ConvertToQuestionAnswerBodyUpdateQuestionAnswerBody ...
func ConvertToQuestionAnswerBodyUpdateQuestionAnswerBody(a QuestionAnswersBodyUpdate) QuestionAnswersBody {
	return QuestionAnswersBody{
		Order:   a.Order,
		Answer:  a.Answer,
		Correct: a.Correct,
	}
}

func (l QuestionBodyUpdate) QuestionNewUpdate() QuestionRaw {
	now := time.Now()
	answers := make([]QuestionAnswersBSON, 0)
	for _, i := range l.Answers {
		answer := i.QuestionNewUpdate()
		answers = append(answers, answer)
	}

	return QuestionRaw{
		ID:        primitive.NewObjectID(),
		Order:     l.Order,
		Question:  l.Question,
		Answers:   answers,
		Active:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// QuestionNewUpdate ...
func (l QuestionAnswersBodyUpdate) QuestionNewUpdate() QuestionAnswersBSON {
	id, _ := primitive.ObjectIDFromHex(l.ID)
	return QuestionAnswersBSON{
		ID:      id,
		Answer:  l.Answer,
		Correct: l.Correct,
		Order:   l.Order,
	}
}
