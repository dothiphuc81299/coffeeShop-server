package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type PackageGroupAdminService struct {
	PackageGroupDAO model.PackageGroupDAO
	GroupDAO        model.GroupDAO
}

// NewPackageGroupAdminService ...
func NewPackageGroupAdminService(d *model.CommonDAO) model.PackageGroupAdminService {
	return &PackageGroupAdminService{
		PackageGroupDAO: d.PackageGroup,
		GroupDAO:        d.Group,
	}
}

func (p *PackageGroupAdminService) Create(ctx context.Context, body model.PackageGroupBody) error {
	payload := body.PackageGroupNewBson()
	if p.checkPackageGroupExisted(ctx, payload.PackageID, payload.GroupID) {
		return errors.New("goi cau hoi da ton tai")
	}

	err := p.PackageGroupDAO.InsertOne(ctx, payload)
	if err != nil {
		fmt.Println("insert package group error", err)
		return err
	}
	return nil
}

func (p *PackageGroupAdminService) checkPackageGroupExisted(ctx context.Context, packageID model.primitive.ObjectID, groupID model.primitive.ObjectID) bool {
	cond := bson.M{
		"packageId": packageID,
		"groupId":   groupID,
	}
	total := p.PackageGroupDAO.CountByCondition(ctx, cond)
	if total >= 1 {
		return true
	}
	return false
}

func (p *PackageGroupAdminService) Update(ctx context.Context, raw model.PackageGroupRaw, body model.PackageGroupBody) error {
	payload := body.PackageGroupNewBson()

	raw.GroupID = payload.GroupID
	raw.PackageID = payload.PackageID
	err := p.PackageGroupDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": raw})

	if err != nil {
		fmt.Println("update package error", err)
		return err
	}
	return nil
}

func (p *PackageGroupAdminService) FindByID(ctx context.Context, id model.primitive.ObjectID) (PackageGroup model.PackageGroupRaw, err error) {
	return p.PackageGroupDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

func (p *PackageGroupAdminService) GetPackageGroupByPackageID(ctx context.Context, packageID model.primitive.ObjectID) []model.PackageGroupAdminResponse {
	var (
		res  = make([]model.PackageGroupAdminResponse, 0)
		cond = bson.M{
			"packageId": packageID,
		}
	)

	docs, _ := p.PackageGroupDAO.FindByCondition(ctx, cond)

	for _, doc := range docs {
		groupRaw, _ := p.GroupDAO.FindOneByCondition(ctx, bson.M{"_id": doc.GroupID})
		quizGroupCommon := model.QuizGroupGetCommon(groupRaw)
		res = append(res, doc.GetPackageGroupAdminResponse(quizGroupCommon))
	}

	return res
}
