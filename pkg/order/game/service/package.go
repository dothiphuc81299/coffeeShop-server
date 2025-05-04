package service

import (
	"context"
	"fmt"
	"sync"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type PackageAdminService struct {
	PackageDAO model.PackageDAO
}

// NewPackageAdminService ...
func NewPackageAdminService(d *model.CommonDAO) model.PackageAdminService {
	return &PackageAdminService{
		PackageDAO: d.Package,
	}
}

func (p *PackageAdminService) Create(ctx context.Context, body model.PackageBody) error {
	payload := body.PackageNewBSON()

	err := p.PackageDAO.InsertOne(ctx, payload)
	if err != nil {
		fmt.Println("insert package error", err)
		return err
	}
	return nil
}

func (p *PackageAdminService) Update(ctx context.Context, raw model.PackageRaw, body model.PackageBody) error {
	payload := body.PackageNewBSON()

	raw.Level = payload.Level
	raw.MinusPoint = payload.MinusPoint
	raw.Reward = payload.Reward
	raw.Title = payload.Title
	raw.UpdatedAt = payload.UpdatedAt
	err := p.PackageDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": raw})

	if err != nil {
		fmt.Println("update package error", err)
		return err
	}
	return nil
}

func (p *PackageAdminService) FindByID(ctx context.Context, id model.primitive.ObjectID) (Group model.PackageRaw, err error) {
	return p.PackageDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (p *PackageAdminService) ListAll(ctx context.Context, q model.CommonQuery) ([]model.PackageAdminResponse, int64) {
	var (
		res   = make([]model.PackageAdminResponse, 0)
		total int64
		cond  = bson.M{}
		wg    sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs, _ := p.PackageDAO.FindByCondition(ctx, cond, q.GetFindOptions())

		for _, doc := range docs {
			res = append(res, model.PackageAdminGetResponse(doc))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total = p.PackageDAO.CountByCondition(ctx, cond)
	}()

	wg.Wait()

	return res, total
}

func (p *PackageAdminService) GetDetail(ctx context.Context, raw model.PackageRaw) model.PackageAdminResponse {
	return model.PackageAdminGetResponse(raw)
}


