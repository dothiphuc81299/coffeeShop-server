package initialize

import (
	"context"
	"fmt"
	"time"

	categoryDAO "github.com/dothiphuc81299/coffeeShop-server/category/dao"
	categoryservice "github.com/dothiphuc81299/coffeeShop-server/category/service"
	drinkDAO "github.com/dothiphuc81299/coffeeShop-server/drink/dao"
	drinkservice "github.com/dothiphuc81299/coffeeShop-server/drink/service"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	roleDAO "github.com/dothiphuc81299/coffeeShop-server/role/dao"
	roleservice "github.com/dothiphuc81299/coffeeShop-server/role/service"

	feedbackDAO "github.com/dothiphuc81299/coffeeShop-server/feedback/dao"
	feedbackservice "github.com/dothiphuc81299/coffeeShop-server/feedback/service"
	orderDAO "github.com/dothiphuc81299/coffeeShop-server/order/dao"
	orderservice "github.com/dothiphuc81299/coffeeShop-server/order/service"
	userDAO "github.com/dothiphuc81299/coffeeShop-server/user/dao"
	userservice "github.com/dothiphuc81299/coffeeShop-server/user/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	eventDAO "github.com/dothiphuc81299/coffeeShop-server/event/dao"

	eventservice "github.com/dothiphuc81299/coffeeShop-server/event/service"

	shiftDAO "github.com/dothiphuc81299/coffeeShop-server/shift/dao"
	shiftservice "github.com/dothiphuc81299/coffeeShop-server/shift/service"

	staffDAO "github.com/dothiphuc81299/coffeeShop-server/staff/dao"
	staffservice "github.com/dothiphuc81299/coffeeShop-server/staff/service"
	staffroleDAO "github.com/dothiphuc81299/coffeeShop-server/staffrole/dao"
	staffroleservice "github.com/dothiphuc81299/coffeeShop-server/staffrole/service"
)

// InitAdminServices ...
func InitAdminServices(d *model.CommonDAO) *model.AdminService {
	return &model.AdminService{
		Drink:    drinkservice.NewDrinkAdminService(d),
		Category: categoryservice.NewCategoryAdminService(d),
		User:     userservice.NewUserAdminService(d),

		Role:     roleservice.NewRoleAdminService(d),
		Order:    orderservice.NewOrderService(d),
		Feedback: feedbackservice.NewFeedbackAdminService(d),

		Event: eventservice.NewEventAdminService(d),

		Shift: shiftservice.NewShiftAdminService(d),

		StaffRole: staffroleservice.NewStaffRoleAdminService(d),
		Staff:     staffservice.NewStaffAdminService(d),
	}
}

// ConnectDB ..
func ConnectDB(dbCfg config.MongoCfg) (*mongo.Database, *model.CommonDAO) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbCfg.URI))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	db := client.Database(dbCfg.Name)

	return db, &model.CommonDAO{
		Drink:    drinkDAO.NewDrinkDAO(db),
		Category: categoryDAO.NewCategoryDAO(db),

		User:     userDAO.NewUserDAO(db),
		Role:     roleDAO.NewRoleDAO(db),
		Order:    orderDAO.NewOrderDAO(db),
		Feedback: feedbackDAO.NewFeedbackDAO(db),

		Event: eventDAO.NewEventDAO(db),

		Shift: shiftDAO.NewShiftDAO(db),

		Staff:     staffDAO.NewStaffDAO(db),
		StaffRole: staffroleDAO.NewStaffRoleDAO(db),
		Session:   staffDAO.NewSessionDAO(db),
	}
}
