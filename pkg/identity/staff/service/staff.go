package service

// import (
// 	"context"
// 	"errors"
// 	"sync"
// 	"time"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // StaffAdminService ...
// type StaffAdminService struct {
// 	StaffDAO   model.StaffDAO
// 	SessionDAO model.SessionDAO
// 	StaffRole  model.StaffRoleDAO
// 	OrderDAO   model.OrderDAO
// }

// // GetToken ...
// func (sfs *StaffAdminService) GetToken(ctx context.Context, staffID model.AppID) (string, error) {
// 	staff, _ := sfs.StaffDAO.FindByID(ctx, staffID)
// 	token := staff.GenerateToken()
// 	// Save session
// 	go func() {
// 		docSession := model.SessionRaw{
// 			ID:        primitive.NewObjectID(),
// 			Staff:     staff.ID,
// 			Token:     util.Base64EncodeToString(token),
// 			CreatedAt: time.Now(),
// 		}
// 		sfs.SessionDAO.InsertOne(context.Background(), docSession)
// 	}()

// 	return token, nil
// }

// // ChangeStatus ...
// func (sfs *StaffAdminService) DeleteStaff(ctx context.Context, data model.Staff) ( err error) {

// 	err = sfs.StaffDAO.DeleteByID(ctx, data.ID)
// 	if err != nil {
// 		return  errors.New(locale.CommonKeyErrorWhenHandle)
// 	}
// 	// Remove session by staff id
// 	go sfs.SessionDAO.RemoveByCondition(context.Background(), bson.M{"staff": data.ID})

// 	return  nil

// }

// // UpdateRole ...
// func (sfs *StaffAdminService) UpdateRole(ctx context.Context, body model.UpdateStaffRoleCommand, data model.Staff) (error) {
// 	roleID, _ := primitive.ObjectIDFromHex(body.Role)
// 	staffRole, err:= sfs.StaffRole.FindByID(ctx, roleID)
// 	if err != nil {
// 		return  err
// 	}

// 	if staffRole.Name == "" {
// 		return  errors.New("StaffRole Is Invalid")
// 	}
// 	// assign
// 	data.Role = roleID
// 	err = sfs.StaffDAO.UpdateByID(ctx, data.ID, bson.M{"$set": data})
// 	if err != nil {
// 		return errors.New(locale.CommonKeyErrorWhenHandle)
// 	}
// 	return  nil
// }

// // FindByID ...
// func (sfs *StaffAdminService) FindByID(ctx context.Context, ID model.AppID) (model.Staff, error) {
// 	return sfs.StaffDAO.FindByID(ctx, ID)
// }

// // Create ...
// func (sfs *StaffAdminService) Create(ctx context.Context, body model.CreateStaffCommand) (res model.StaffGetResponseAdmin, err error) {
// 	// Check username staff existed
// 	isExisted := sfs.checkUserExisted(ctx, body.Username)
// 	if isExisted {
// 		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonyKeyUserNameIsExisted)
// 	}

// 	roleID, _ := primitive.ObjectIDFromHex(body.Role)
// 	staffRole, _ := sfs.StaffRole.FindByID(ctx, roleID)

// 	// Create
// 	doc := body.StaffNewBSON(staffRole.Permissions)
// 	err = sfs.StaffDAO.InsertOne(ctx, doc)

// 	if err != nil {
// 		return
// 	}
// 	return doc.GetStaffResponseAdmin(), err
// }

// func (sfs *StaffAdminService) checkUserExisted(ctx context.Context, username string) bool {
// 	total := sfs.StaffDAO.CountByCondition(ctx, bson.M{"username": username})
// 	return total > 0
// }

// // ListStaff ...
// func (sfs *StaffAdminService) ListStaff(ctx context.Context, q model.CommonQuery) ([]model.StaffGetResponseAdmin, int64) {
// 	var (
// 		wg   sync.WaitGroup
// 		res  = make([]model.StaffGetResponseAdmin, 0)
// 		cond = bson.M{
// 			"isRoot": false,
// 		}
// 		total int64
// 	)
// 	q.AssignActive(&cond)
// 	q.AssignUsername(&cond)

// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		docs, _ := sfs.StaffDAO.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
// 		for _, s := range docs {
// 			staff := s.GetStaffResponseAdmin()
// 			res = append(res, staff)
// 		}
// 	}()
// 	go func() {
// 		defer wg.Done()
// 		total = sfs.StaffDAO.CountByCondition(ctx, cond)
// 	}()
// 	wg.Wait()
// 	return res, total
// }

// func (sfs *StaffAdminService) GetDetailStaff(ctx context.Context, staff model.Staff) model.StaffMeResponse {
// 	return model.StaffMeResponse{
// 		ID:          staff.ID,
// 		Username:    staff.Username,
// 		Phone:       staff.Phone,
// 		Avatar:      staff.Avatar,
// 		Permissions: staff.Permissions,
// 		Address:     staff.Address,
// 		Token:       staff.GenerateToken(),
// 	}
// }

// func (sfs *StaffAdminService) LoginStaff(ctx context.Context, body model.LoginStaffCommand) (model.StaffResponse, error) {
// 	cond := bson.M{
// 		"username": body.Username,
// 		"password": body.Password,
// 	}

// 	staff, err := sfs.StaffDAO.FindOneByCondition(ctx, cond)
// 	if err != nil {
// 		return model.StaffResponse{}, errors.New(locale.UserNameOrPasswordIsIncorrect)
// 	}
// 	if !staff.Active {
// 		return model.StaffResponse{},errors.New(locale.CommonKeyStaffIsDeleted)
// 	}
// 	token, err := sfs.GetToken(ctx, staff.ID)
// 	if err != nil {
// 		return model.StaffResponse{}, err
// 	}
// 	doc := staff.GetStaffResponse(token)
// 	return doc, nil
// }

// func (sfs *StaffAdminService) GetStaffByID(ctx context.Context, id model.AppID) model.StaffGetResponseAdmin {
// 	staff, err := sfs.FindByID(ctx, id)
// 	if err != nil {
// 		return model.StaffGetResponseAdmin{}
// 	}
// 	return  staff.GetStaffResponseAdmin()
// }

// // NewStaffAdminService ...
// func NewStaffAdminService(sd *model.CommonDAO) model.StaffAdminService {
// 	return &StaffAdminService{
// 		StaffDAO:   sd.Staff,
// 		SessionDAO: sd.Session,
// 		StaffRole:  sd.StaffRole,
// 		OrderDAO:   sd.Order,
// 	}
// }
