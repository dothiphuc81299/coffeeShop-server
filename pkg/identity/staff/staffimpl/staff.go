package staffimpl

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	store     *store
	roleStore role.Store
}

func NewService(store *store, roleStore role.Store) staff.Service {
	return &service{
		store:     store,
		roleStore: roleStore,
	}
}

func (s *service) Create(ctx context.Context, cmd staff.CreateStaffCommand) error {
	isExisted := s.checkUserExisted(ctx, cmd.Username)
	if isExisted {
		return staff.ErrStaffExisted
	}

	roleID, _ := primitive.ObjectIDFromHex(cmd.Role)
	roleStore, _ := s.roleStore.FindByID(ctx, roleID)

	err := s.store.InsertOne(ctx, staff.Staff{
		ID:          primitive.NewObjectID(),
		Username:    cmd.Username,
		Password:    cmd.Password,
		Role:        roleID,
		Active:      true,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Permissions: roleStore.Permissions,
		Address:     cmd.Address,
		Phone:       cmd.Phone,
		IsRoot:      false,
	})

	if err != nil {
		return err
	}
	return err
}

func (s *service) Update(ctx context.Context, cmd staff.UpdateStaffCommand, data staff.Staff) error {
	payload := bson.M{
		"address":   cmd.Address,
		"phone":     cmd.Phone,
		"updatedAt": time.Now().UTC(),
	}

	err := s.store.UpdateByID(ctx, data.ID, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChangePassword(ctx context.Context, staff staff.Staff, body staff.PasswordBody) error {
	res, _ := s.store.FindOneByCondition(ctx, bson.M{"_id": staff.ID})
	if res.ID.IsZero() {
		return errors.New(locale.CommonKeyStaffIsDeleted)
	}

	if body.Password != res.Password || body.NewPassword != body.NewPasswordAgain || body.NewPassword == body.Password {
		return errors.New(locale.PasswordIsIncorrect)
	}

	err := s.store.UpdateByID(ctx, staff.ID, bson.M{"$set": bson.M{"password": body.NewPassword}})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) FindByID(ctx context.Context, ID primitive.ObjectID) (staff.Staff, error) {
	return s.store.FindByID(ctx, ID)
}

func (s *service) DeleteStaff(ctx context.Context, data staff.Staff) (err error) {
	err = s.store.DeleteByID(ctx, data.ID)
	if err != nil {
		return errors.New(locale.CommonKeyErrorWhenHandle)
	}

	// TODO :remove token
	return nil
}

// UpdateRole ...
func (s *service) UpdateRole(ctx context.Context, body staff.UpdateStaffRoleCommand, data staff.Staff) error {
	roleID, _ := primitive.ObjectIDFromHex(body.Role)
	roleStore, err := s.roleStore.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	if roleStore.Name == "" {
		return errors.New("roleStore Is Invalid")
	}

	data.Role = roleID
	err = s.store.UpdateByID(ctx, data.ID, bson.M{"$set": data})
	if err != nil {
		return errors.New(locale.CommonKeyErrorWhenHandle)
	}
	return nil
}

func (s *service) FindByID(ctx context.Context, ID staff.AppID) (staff.Staff, error) {
	return s.store.FindByID(ctx, ID)
}

func (s *service) checkUserExisted(ctx context.Context, username string) bool {
	total := s.store.CountByCondition(ctx, bson.M{"username": username})
	return total > 0
}

// ListStaff ...
func (s *service) ListStaff(ctx context.Context, q query.CommonQuery) ([]staff.StaffGetResponseAdmin, int64) {
	var (
		wg   sync.WaitGroup
		res  = make([]staff.StaffGetResponseAdmin, 0)
		cond = bson.M{
			"isRoot": false,
		}
		total int64
	)
	q.AssignActive(&cond)
	q.AssignUsername(&cond)

	wg.Add(2)
	go func() {
		defer wg.Done()
		docs, _ := s.store.FindByCondition(ctx, cond, q.GetFindOptsUsingPage())
		for _, s := range docs {
			staff := s.GetStaffResponseAdmin()
			res = append(res, staff)
		}
	}()
	go func() {
		defer wg.Done()
		total = s.store.CountByCondition(ctx, cond)
	}()
	wg.Wait()
	return res, total
}

func (s *service) GetDetailStaff(ctx context.Context, staff staff.Staff) staff.StaffMeResponse {
	return staff.StaffMeResponse{
		ID:          staff.ID,
		Username:    staff.Username,
		Phone:       staff.Phone,
		Avatar:      staff.Avatar,
		Permissions: staff.Permissions,
		Address:     staff.Address,
		Token:       staff.GenerateToken(),
	}
}

func (s *service) LoginStaff(ctx context.Context, body staff.LoginStaffCommand) (staff.StaffResponse, error) {
	cond := bson.M{
		"username": body.Username,
		"password": body.Password,
	}

	staff, err := s.store.FindOneByCondition(ctx, cond)
	if err != nil {
		return staff.StaffResponse{}, errors.New(locale.UserNameOrPasswordIsIncorrect)
	}
	if !staff.Active {
		return staff.StaffResponse{}, errors.New(locale.CommonKeyStaffIsDeleted)
	}
	token, err := s.GetToken(ctx, staff.ID)
	if err != nil {
		return staff.StaffResponse{}, err
	}
	doc := staff.GetStaffResponse(token)
	return doc, nil
}

func (s *service) GetStaffByID(ctx context.Context, id staff.AppID) staff.StaffGetResponseAdmin {
	staff, err := s.FindByID(ctx, id)
	if err != nil {
		return staff.StaffGetResponseAdmin{}
	}
	return staff.GetStaffResponseAdmin()
}
