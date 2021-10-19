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
	StaffRole  model.StaffRoleDAO
	OrderDAO   model.OrderDAO
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
	roleID, _ := primitive.ObjectIDFromHex(body.Role)
	staffRole, _ := sfs.StaffRole.FindByID(ctx, roleID)

	doc := body.StaffNewBSON(staffRole.Permissions)

	// assign
	data.Address = doc.Address
	data.Permissions = doc.Permissions
	data.Username = doc.Username
	data.Phone = doc.Phone
	data.Role = doc.Role
	err := sfs.StaffDAO.UpdateByID(ctx, data.ID, bson.M{"$set": data})
	if err != nil {
		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonKeyErrorWhenHandle)
	}

	return data.GetStaffResponseAdmin(), nil
}

// FindByID ...
func (sfs *StaffAdminService) FindByID(ctx context.Context, ID model.AppID) (model.StaffRaw, error) {
	return sfs.StaffDAO.FindByID(ctx, ID)
}

// Create ...
func (sfs *StaffAdminService) Create(ctx context.Context, body model.StaffBody) (res model.StaffGetResponseAdmin, err error) {
	// Check username staff existed
	isExisted := sfs.checkUserExisted(ctx, body.Username)
	if isExisted {
		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonKeyPhoneExisted)
	}

	roleID, _ := primitive.ObjectIDFromHex(body.Role)
	staffRole, _ := sfs.StaffRole.FindByID(ctx, roleID)

	// Create
	doc := body.StaffNewBSON(staffRole.Permissions)
	err = sfs.StaffDAO.InsertOne(ctx, doc)

	if err != nil {
		return
	}
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
		docs, _ := sfs.StaffDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
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

func (sfs *StaffAdminService) GetDetailStaff(ctx context.Context, staff model.StaffRaw) model.StaffMeResponse {
	return model.StaffMeResponse{
		ID:          staff.ID,
		Username:    staff.Username,
		Phone:       staff.Phone,
		Avatar:      staff.Avatar,
		Permissions: staff.Permissions,
		Address:     staff.Address,
		Token:       staff.GenerateToken(),
	}
}

func (sfs *StaffAdminService) StaffLogin(ctx context.Context, body model.StaffLoginBody) (model.StaffResponse, error) {
	cond := bson.M{
		"username": body.Username,
		"password": body.Password,
	}

	staff, err := sfs.StaffDAO.FindOneByCondition(ctx, cond)
	if err != nil {
		return model.StaffResponse{}, err
	}
	token, err := sfs.GetToken(ctx, staff.ID)
	if err != nil {
		return model.StaffResponse{}, err
	}
	doc := staff.GetStaffResponse(token)
	return doc, nil
}

func (sfs *StaffAdminService) getMonthNow(cond bson.M, date string) {
	now := time.Now()
	y, m, _ := now.Date()

	from := time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
	to := time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
	to = util.TimeStartOfDayInHCM(to)
	from = util.TimeStartOfDayInHCM(from)
	cond[date] = bson.M{
		"$gte": from,
		"$lt":  to,
	}
}

// NewStaffAdminService ...
func NewStaffAdminService(sd *model.CommonDAO) model.StaffAdminService {
	return &StaffAdminService{
		StaffDAO:   sd.Staff,
		SessionDAO: sd.Session,
		StaffRole:  sd.StaffRole,
		OrderDAO:   sd.Order,
	}
}
