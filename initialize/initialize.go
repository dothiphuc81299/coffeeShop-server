package initialize

import (
	drinkservice "github.com/dothiphuc81299/coffeeShop-server/drink/service"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
)

// InitAdminService ...
func InitAdminService(d *model.CommonDAO) *model.AdminService {
	return &model.AdminService{
		Drink: drinkservice.NewDrinkAdminService(d),
	}
}
