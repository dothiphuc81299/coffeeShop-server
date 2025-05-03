package service

import (
	"context"

	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
)

type UserPackageAppService struct {
	UserPackage model.UserPackageGroupDAO
}

func NewUserPackageAppService(d *model.CommonDAO) model.UserPackageGroupAppService {
	return &UserPackageAppService{
		UserPackage: d.UserPackageGroup,
	}
}

func (u *UserPackageAppService) ChoosePakage(ctx context.Context, body model.UserPackageBody) error {
	panic("TODO")
}
