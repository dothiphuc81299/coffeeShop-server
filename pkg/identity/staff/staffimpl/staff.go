package staffimpl

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/staff/role"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/identity/token"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/password"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/util/query"
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

	password, err := password.HashPassword(cmd.Password)
	if err != nil {
		return err
	}

	err = s.store.InsertOne(ctx, staff.Staff{
		ID:          primitive.NewObjectID(),
		Username:    cmd.Username,
		Password:    password,
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

func (s *service) Update(ctx context.Context, cmd *staff.UpdateStaffCommand) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.Staff {
		return errors.New("account is invalid")
	}

	payload := bson.M{
		"address":   cmd.Address,
		"phone":     cmd.Phone,
		"updatedAt": time.Now().UTC(),
	}

	err := s.store.UpdateByID(ctx, account.ID, bson.M{"$set": payload})
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChangePassword(ctx context.Context, body *staff.PasswordBody) error {
	account, ok := ctx.Value("current_account").(*token.AccountData)
	if !ok || account.AccountType != token.Staff {
		return errors.New("account is invalid")
	}

	res, _ := s.store.FindOneByCondition(ctx, bson.M{"_id": account.ID})
	if res.ID.IsZero() {
		return staff.ErrStaffNotFound
	}

	if !password.CheckPassword(body.Password, res.Password) {
		return staff.ErrPasswordInvalid
	}

	if body.NewPassword != body.NewPasswordAgain || body.NewPassword == body.Password {
		return staff.ErrPasswordInvalid
	}

	hashPassword, err := password.HashPassword(body.NewPassword)
	if err != nil {
		return err
	}

	err = s.store.UpdateByID(ctx, account.ID, bson.M{"$set": bson.M{"password": hashPassword}})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateRole(ctx context.Context, cmd *staff.UpdateStaffRoleCommand) error {
	staffResult, err := s.store.FindByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if staffResult.ID.IsZero() {
		return staff.ErrStaffNotFound
	}

	roleID, _ := primitive.ObjectIDFromHex(cmd.Role)
	roleStore, err := s.roleStore.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	if roleStore.Name == "" {
		return errors.New("roleStore Is Invalid")
	}

	err = s.store.UpdateByID(ctx, cmd.ID, bson.M{"$set": bson.M{"role": roleID, "permissions": roleStore.Permissions, "updatedAt": time.Now().UTC()}})
	if err != nil {
		return staff.ErrCanNotUpdateRole
	}
	return nil
}

func (s *service) checkUserExisted(ctx context.Context, username string) bool {
	total := s.store.CountByCondition(ctx, bson.M{"username": username})
	return total > 0
}

func (s *service) ListStaff(ctx context.Context, q *query.CommonQuery) ([]staff.Staff, int64) {
	var (
		wg   sync.WaitGroup
		res  = make([]staff.Staff, 0)
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
		res = docs
	}()
	go func() {
		defer wg.Done()
		total = s.store.CountByCondition(ctx, cond)
	}()
	wg.Wait()
	return res, total
}

func (s *service) LoginStaff(ctx context.Context, body staff.LoginStaffCommand) (*staff.StaffResponse, error) {
	cond := bson.M{
		"username": body.Username,
	}

	entity, err := s.store.FindOneByCondition(ctx, cond)

	if err != nil {
		return nil, staff.ErrUserNameOrPasswordIsIncorrect
	}
	if !entity.Active {
		return nil, staff.ErrStaffIsDeleted
	}

	if !password.CheckPassword(body.Password, entity.Password) {
		return nil, staff.ErrUserNameOrPasswordIsIncorrect
	}

	var role token.AccountType
	if entity.IsRoot {
		role = token.Root
	} else {
		role = token.Staff
	}

	tokenStr, err := token.GenerateJWT(entity.ID, entity.Username, role)
	if err != nil {
		return nil, err
	}

	return &staff.StaffResponse{
		ID:          entity.ID,
		Username:    entity.Username,
		Phone:       entity.Phone,
		Address:     entity.Address,
		Token:       tokenStr,
		Permissions: entity.Permissions,
	}, nil
}

func (s *service) GetStaffByID(ctx context.Context, id primitive.ObjectID) (staff.Staff, error) {
	result, err := s.store.FindByID(ctx, id)
	if err != nil {
		return staff.Staff{}, err
	}

	result.Password = ""
	return result, nil
}
