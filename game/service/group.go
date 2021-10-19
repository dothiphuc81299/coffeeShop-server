package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GroupAdminService ...
type GroupAdminService struct {
	GroupDAO model.GroupDAO
}

// NewGroupAdminService ...
func NewGroupAdminService(d *model.CommonDAO) model.GroupAdminService {
	return &GroupAdminService{
		GroupDAO: d.Group,
	}
}

func (g *GroupAdminService) Create(ctx context.Context, body model.QuizGroupBody) error {
	payload := body.QuizGroupNewBSON()
	for _, value := range payload.Quizzes {
		if g.checkQuizExistedByID(ctx, value) != 1 {
			return errors.New("QuizKeyQuizIDNotExisted")
		}
	}

	err := g.GroupDAO.InsertOne(ctx, payload)
	if err != nil {
		fmt.Println("insert quiz group error", err)
		return err
	}
	return nil
}

func (g *GroupAdminService) checkQuizExistedByID(ctx context.Context, id primitive.ObjectID) int {
	total := g.GroupDAO.CountByCondition(ctx, bson.M{"_id": id})
	return int(total)
}

func (g *GroupAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.QuizGroupCommon, int64) {
	var (
		res   = make([]model.QuizGroupCommon, 0)
		total int64
		cond  = bson.M{}
		wg    sync.WaitGroup
	)

	q.AssignActive(&cond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs, _ := g.GroupDAO.FindByCondition(ctx, cond, q.GetFindOptions())

		for _, doc := range docs {
			res = append(res, model.QuizGroupGetCommon(doc))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total = g.GroupDAO.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	return res, total
}

func (g *GroupAdminService) Update(ctx context.Context, group model.QuizGroupRaw, body model.QuizGroupBody) error {
	payload := body.QuizGroupNewBSON()
	for _, value := range payload.Quizzes {
		if g.checkQuizExistedByID(ctx, value) != 1 {
			return errors.New("LuckyDrawQuizKeyQuizIDNotExisted")
		}
	}

	group.Quizzes = payload.Quizzes
	group.Name = payload.Name
	group.UpdatedAt = payload.UpdatedAt

	err := g.GroupDAO.UpdateByID(ctx, group.ID, bson.M{"$set": group})

	if err != nil {
		fmt.Println("update quiz group error", err)
		return err
	}
	return nil
}

func (g *GroupAdminService) FindByID(ctx context.Context, id model.AppID) (Group model.QuizGroupRaw, err error) {
	return g.GroupDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (g *GroupAdminService) ChangeStatus(ctx context.Context, group model.QuizGroupRaw) (status bool, err error) {
	active := !group.Active
	payload := bson.M{
		"active":    active,
		"updatedAt": time.Now(),
	}

	err = g.GroupDAO.UpdateByID(ctx, group.ID, bson.M{"$set": payload})
	if err != nil {
		fmt.Println("update quiz group error", err)
		return
	}
	return active, nil
}
