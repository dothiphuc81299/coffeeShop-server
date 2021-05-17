package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type StaffAppService struct {
	StaffDAO  model.StaffDAO
	StaffRole model.StaffRoleDAO
}

func NewStaffAppService(d *model.CommonDAO) model.StaffAppService {
	return &StaffAppService{
		StaffDAO:  d.Staff,
		StaffRole: d.StaffRole,
	}
}

// Update ...
func (sfs *StaffAppService) Update(ctx context.Context, body model.StaffUpdateBodyByIt, data model.StaffRaw) (model.StaffGetResponseAdmin, error) {
	log.Println("7")
	payload := bson.M{
		"username":  body.Username,
		"address":   body.Address,
		"phone":     body.Phone,
		"updatedAt": time.Now(),
	}

	log.Println("8")
	err := sfs.StaffDAO.UpdateByID(ctx, data.ID, bson.M{"$set": payload})
	log.Println("9")
	if err != nil {
		return model.StaffGetResponseAdmin{}, errors.New(locale.CommonKeyErrorWhenHandle)
	}
	log.Println("10")
	return data.GetStaffResponseAdmin(), nil
}

func (sfs *StaffAppService) ChangePassword(ctx context.Context, staff model.StaffRaw, body model.PasswordBody) (err error) {
	res, _ := sfs.StaffDAO.FindOneByCondition(ctx, bson.M{"_id": staff.ID})
	if res.ID.IsZero() {
		return errors.New("staff khong ton tai")
	}

	if body.Password != res.Password || body.NewPassword != body.NewPasswordAgain || body.NewPassword == body.Password {
		return errors.New("mat khau  khong dung")
	}

	err = sfs.StaffDAO.UpdateByID(ctx, staff.ID, bson.M{"$set": bson.M{"password": body.NewPassword}})
	if err != nil {
		return
	}
	return
}

// FindByID ...
func (sfs *StaffAppService) FindByID(ctx context.Context, ID model.AppID) (model.StaffRaw, error) {
	return sfs.StaffDAO.FindByID(ctx, ID)
}
