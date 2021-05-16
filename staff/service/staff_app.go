package service

import (
	"context"
	"errors"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (sfs *StaffAppService) Update(ctx context.Context, body model.StaffBody, data model.StaffRaw) (model.StaffGetResponseAdmin, error) {
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
