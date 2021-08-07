package service

import (
	"context"
	"errors"

	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// RoleAdminService ...
type RoleAdminService struct {
	RoleDAO model.RoleDAO
}

// Update ...
func (rs *RoleAdminService) Update(ctx context.Context, body model.RoleBody, raw model.RoleRaw) (doc model.RoleAdminResponse, err error) {
	raw.Name = body.Name
	raw.Permissions = body.Permissions
	raw.UpdatedAt = time.Now()
	err = rs.RoleDAO.UpdateByID(ctx, raw.ID, bson.M{"$set": raw})
	if err != nil {
		return doc, errors.New(locale.CommonKeyErrorWhenHandle)
	}

	docBson, _ := rs.RoleDAO.FindOneByCondition(ctx, bson.M{"_id": raw.ID})
	return docBson.GetResponse(), nil
}

// FindByID ...
func (rs *RoleAdminService) FindByID(ctx context.Context, id model.AppID) (model.RoleRaw, error) {
	return rs.RoleDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

// List ...
func (rs *RoleAdminService) List(ctx context.Context, q model.CommonQuery) ([]model.RoleAdminResponse, int64) {
	var (
		cond  = bson.M{}
		res   = make([]model.RoleAdminResponse, 0)
		wg    sync.WaitGroup
		total int64
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		docs, _ := rs.RoleDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
		for _, doc := range docs {
			item := model.RoleAdminResponse{
				ID:          doc.ID,
				Name:        doc.Name,
				Permissions: doc.Permissions,
				CreatedAt:   doc.CreatedAt,
			}
			res = append(res, item)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total = rs.RoleDAO.CountByCondition(ctx, cond)

	}()
	wg.Wait()

	return res, total
}

// Create ...
func (rs *RoleAdminService) Create(ctx context.Context, body model.RoleBody) (doc model.RoleAdminResponse, err error) {
	now := time.Now()
	roleNewBSON := model.RoleRaw{
		ID:          model.NewAppID(),
		Name:        body.Name,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: body.Permissions,
	}
	err = rs.RoleDAO.InsertOne(ctx, roleNewBSON)

	if err != nil {
		return doc, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	return roleNewBSON.GetResponse(), nil
}

func (rs *RoleAdminService) GetDetail(ctx context.Context, role model.RoleRaw) model.RoleAdminResponse {
	return role.GetResponse()
}

// NewRoleAdminService ...
func NewRoleAdminService(d *model.CommonDAO) model.RoleService {
	return &RoleAdminService{
		RoleDAO: d.Role,
	}
}
