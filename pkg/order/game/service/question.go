package service

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
)

// QuestionAdminService ...
type QuestionAdminService struct {
	QuestionDAO model.QuestionDAO
}

// NewQuestionAdminService ...
func NewQuestionAdminService(d *model.CommonDAO) model.QuestionAdminService {
	return &QuestionAdminService{
		QuestionDAO: d.Question,
	}
}

func (q *QuestionAdminService) Create(ctx context.Context, body model.QuestionBody) error {
	payload := body.QuestionNewBSON()
	err := checkNumberOfCorrectAnswers(body.Answers)
	if err != nil {
		return err
	}

	err = q.QuestionDAO.InsertOne(ctx, payload)
	if err != nil {
		fmt.Println("insert quiz error", err)
		return err
	}
	return nil
}

func checkNumberOfCorrectAnswers(c []model.QuestionAnswersBody) (err error) {
	a := funk.Filter(c, func(ans model.QuestionAnswersBody) bool {
		return ans.Correct == true
	})

	if a == nil || reflect.ValueOf(a).Len() != 1 {
		return errors.New("Answer Correct Must Have One True")
	}
	return

}

func (q *QuestionAdminService) ListAll(ctx context.Context, query model.CommonQuery) ([]model.QuestionCommon, int64) {
	var (
		result = make([]model.QuestionCommon, 0)
		total  int64
		wg     sync.WaitGroup
		cond   = bson.M{}
	)

	// Assign condition
	query.AssignActive(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		quizzes, _ := q.QuestionDAO.FindByCondition(ctx, cond, query.GetFindOptions())
		for _, value := range quizzes {
			temp := model.QuestionGetCommonAPI(value)
			result = append(result, temp)
		}
	}()
	go func() {
		defer wg.Done()
		total = q.QuestionDAO.CountByCondition(ctx, cond)
	}()

	wg.Wait()
	return result, total
}

func (q *QuestionAdminService) Update(ctx context.Context, quiz model.QuestionRaw, body model.QuestionBodyUpdate) error {
	ans := make([]model.QuestionAnswersBody, 0)
	for _, value := range body.Answers {
		ansBody := model.ConvertToQuestionAnswerBodyUpdateQuestionAnswerBody(value)
		ans = append(ans, ansBody)
	}

	err := checkNumberOfCorrectAnswers(ans)
	if err != nil {
		return err
	}

	payload := body.QuestionNewUpdate()
	quiz.Order = payload.Order
	quiz.UpdatedAt = payload.UpdatedAt
	quiz.Question = payload.Question
	quiz.Answers = payload.Answers

	err = q.QuestionDAO.UpdateByID(ctx, quiz.ID, bson.M{"$set": quiz})

	if err != nil {
		fmt.Println("update quiz error", err)
		return err
	}
	return nil
}

func (q *QuestionAdminService) FindByID(ctx context.Context, id model.AppID) (Question model.QuestionRaw, err error) {
	return q.QuestionDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (q *QuestionAdminService) GetDetail(ctx context.Context, quest model.QuestionRaw) model.QuestionCommon {
	return model.QuestionGetCommonAPI(quest)
}

func (q *QuestionAdminService) ChangeStatus(ctx context.Context, quiz model.QuestionRaw) (status bool, err error) {
	active := !quiz.Active
	payload := bson.M{
		"active":    active,
		"updatedAt": time.Now(),
	}

	err = q.QuestionDAO.UpdateByID(ctx, quiz.ID, bson.M{"$set": payload})
	if err != nil {
		fmt.Println("change status quiz error", err)
		return
	}
	return active, err
}
