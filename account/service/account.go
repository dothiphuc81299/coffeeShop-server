package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// AccountAdminService ...
type AccountAdminService struct {
	AccountDAO model.AccountDAO
	UserDAO    model.UserDAO
}

// NewAccountAdminService ...
func NewAccountAdminService(d *model.CommonDAO) model.AccountAdminService {
	return &AccountAdminService{
		AccountDAO: d.Account,
		UserDAO:    d.User,
	}
}

// GenerateToken ...
func (as *AccountAdminService) GenerateToken(ctx context.Context, acc model.AccountRaw, userID model.AppID) (string, error) {
	// Check account active
	if !acc.Active {
		return "", errors.New(locale.AuthKeyAccountUnActive)
	}

	// Check user
	user := new(model.UserRaw)
	as.UserDAO.FindOneByCondition(ctx, bson.M{"user": userID})
	if user.ID.IsZero() || !user.Active {
		return "", errors.New(locale.AuthKeyUserUnActive)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":   acc.ID,
		"phone": user.Phone,
		"user":  userID,
		"exp":   time.Now().Local().Add(time.Second * 15552000).Unix(), // 6 months
	})
	tokenString, _ := token.SignedString([]byte(config.GetEnv().AuthSecret))

	return tokenString, nil
}

// FindByID ...
func (as *AccountAdminService) FindByID(ctx context.Context, id model.AppID) (model.AccountRaw, error) {
	return as.AccountDAO.FindOneByCondition(ctx, bson.M{"_id": id})
}

// Update ...
func (as *AccountAdminService) Update(ctx context.Context, body model.AccountBody, raw model.AccountRaw) (model.AccountResponse, error) {

	raw.Permissions = body.Permissions
	raw.UpdatedAt = time.Now()
	err := as.AccountDAO.UpdateByID(ctx, raw.ID, &raw)
	doc, _ := as.AccountDAO.FindOneByCondition(ctx, bson.M{"_id": raw.ID})

	return doc.GetResponse(), err
}

// ChangeStatus ...
func (as *AccountAdminService) ChangeStatus(ctx context.Context, raw model.AccountRaw) (active bool, err error) {
	raw.Active = !raw.Active
	raw.UpdatedAt = time.Now()
	err = as.AccountDAO.UpdateByID(ctx, raw.ID, &raw)
	if err != nil {
		return
	}
	return raw.Active, err
}
