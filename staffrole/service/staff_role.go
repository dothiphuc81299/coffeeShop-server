package service

import (
	"context"
	"errors"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StaffRoleAdminService ...
type StaffRoleAdminService struct {
	StaffRoleDAO model.StaffRoleDAO
}

// FindByID ...
func (ss *StaffRoleAdminService) FindByID(ctx context.Context, id model.AppID) (model.StaffRoleRaw, error) {
	return ss.StaffRoleDAO.FindByID(ctx, id)
}

// Update ...
func (ss *StaffRoleAdminService) Update(ctx context.Context, data model.StaffRoleRaw, body model.StaffRoleBody) (model.StaffRoleAdminResponse, error) {
	payload := bson.M{
		"$set": bson.M{
			"name":        body.Name,
			"permissions": body.Permissions,
			"updatedAt":   time.Now(),
		},
	}

	err := ss.StaffRoleDAO.UpdateByID(ctx, data.ID, payload)
	if err != nil {
		return model.StaffRoleAdminResponse{}, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	data.Name = body.Name
	data.Permissions = body.Permissions
	return data.GetResponse(), nil
}

// ListStaffRole ...
func (ss *StaffRoleAdminService) ListStaffRole(ctx context.Context, q model.CommonQuery) ([]model.StaffRoleAdminResponse, int64) {
	res := make([]model.StaffRoleAdminResponse, 0)
	cond := bson.M{}
	docs, _ := ss.StaffRoleDAO.FindByCondition(ctx, cond)
	for _, doc := range docs {
		item := model.StaffRoleAdminResponse{
			ID:          doc.ID,
			Name:        doc.Name,
			Permissions: doc.Permissions,
			CreatedAt:   model.TimeResponseInit(doc.CreatedAt),
		}
		res = append(res, item)
	}
	total := ss.StaffRoleDAO.CountByCondition(ctx, cond)
	return res, total
}

// Create ...
func (ss *StaffRoleAdminService) Create(ctx context.Context, body model.StaffRoleBody) (model.StaffRoleAdminResponse, error) {
	now := time.Now()
	roleNewBSON := model.StaffRoleRaw{
		ID:          primitive.NewObjectID(),
		Name:        body.Name,
		CreatedAt:   now,
		UpdatedAt:   now,
		Permissions: body.Permissions,
	}

	err := ss.StaffRoleDAO.InsertOne(ctx, &roleNewBSON)
	if err != nil {
		return model.StaffRoleAdminResponse{}, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	return roleNewBSON.GetResponse(), nil
}

// NewStaffRoleAdminService ...
func NewStaffRoleAdminService(srd model.StaffRoleDAO) model.StaffRoleAdminService {
	return &StaffRoleAdminService{StaffRoleDAO: srd}
}
