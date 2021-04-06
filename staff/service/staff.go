package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StaffAdminService ...
type StaffAdminService struct {
	StaffDAO   model.StaffDAO
	SessionDAO model.SessionDAO
}

// GetToken ...
func (sfs *StaffAdminService) GetToken(ctx context.Context, staffID model.AppID) (string, error) {
	staff, _ := sfs.StaffDAO.FindByID(ctx, staffID)
	token := staff.GenerateToken()
	// Save session
	go func() {
		docSession := model.SessionRaw{
			ID:        primitive.NewObjectID(),
			Staff:     staff.ID,
			Token:     util.Base64EncodeToString(token),
			CreatedAt: time.Now(),
		}
		sfs.SessionDAO.InsertOne(context.Background(), docSession)
	}()

	return token, nil
}

// ChangeStatus ...
func (sfs *StaffAdminService) ChangeStatus(ctx context.Context, data model.StaffRaw) (active bool, err error) {
	payload := bson.M{
		"$set": bson.M{
			"active":    !data.Active,
			"updatedAt": time.Now(),
		},
	}
	err = sfs.StaffDAO.UpdateByID(ctx, data.ID, payload)
	if err != nil {
		return active, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	data.Active = !data.Active

	// Remove session by staff id
	go sfs.SessionDAO.RemoveByCondition(context.Background(), bson.M{"staff": data.ID})

	return data.Active, nil

}

// Update ...
func (sfs *StaffAdminService) Update(ctx context.Context, body model.StaffBody, data model.StaffRaw) (model.StaffGetResponseAdmin, error) {
	payload := bson.M{
		"$set": bson.M{
			"username":    body.Username,
			"phone":       body.Phone,
			"address":     body.Address,
			"avatar":      body.Avatar,
			"permissions": body.Permissions,
			"updatedAt":   time.Now(),
		},
	}

	err := sfs.StaffDAO.UpdateByID(ctx, data.ID, payload)
	if err != nil {
		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	data.Username = body.Username
	data.Permissions = body.Permissions

	// Remove session by staff id
	go sfs.SessionDAO.RemoveByCondition(context.Background(), bson.M{"staff": data.ID})

	return data.GetStaffResponseAdmin(), nil
}

// FindByID ...
func (sfs *StaffAdminService) FindByID(ctx context.Context, ID model.AppID) (model.StaffRaw, error) {
	return sfs.StaffDAO.FindByID(ctx, ID)
}

// Create ...
func (sfs *StaffAdminService) Create(ctx context.Context, body model.StaffBody) (model.StaffGetResponseAdmin, error) {
	// Check username staff existed
	isExisted := sfs.checkUserExisted(ctx, body.Username)
	if isExisted {
		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonKeyPhoneExisted)
	}
	// Create
	doc := body.StaffNewBSON()
	err := sfs.StaffDAO.InsertOne(ctx, doc)

	return doc.GetStaffResponseAdmin(), err
}

func (sfs *StaffAdminService) checkUserExisted(ctx context.Context, username string) bool {
	total := sfs.StaffDAO.CountByCondition(ctx, bson.M{"username": username})
	return total > 0
}

// ListStaff ...
func (sfs *StaffAdminService) ListStaff(ctx context.Context, q model.CommonQuery) ([]model.StaffGetResponseAdmin, int64) {
	var (
		wg    sync.WaitGroup
		res   = make([]model.StaffGetResponseAdmin, 0)
		cond  = bson.M{}
		total int64
	)
	q.AssignActive(&cond)
	q.AssignKeyword(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		docs, _ := sfs.StaffDAO.FindByCondition(ctx, cond)
		for _, s := range docs {
			staff := s.GetStaffResponseAdmin()
			res = append(res, staff)
		}
	}()
	go func() {
		defer wg.Done()
		total = sfs.StaffDAO.CountByCondition(ctx, cond)
	}()
	wg.Wait()
	return res, total
}

// NewStaffAdminService ...
func NewStaffAdminService(sd *model.CommonDAO) model.StaffAdminService {
	return &StaffAdminService{
		StaffDAO:   sd.Staff,
		SessionDAO: sd.Session,
	}
}