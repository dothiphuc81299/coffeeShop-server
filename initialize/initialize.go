package initialize

import (
	"context"
	"fmt"
	"time"

	accountDAO "github.com/dothiphuc81299/coffeeShop-server/account/dao"
	accountservice "github.com/dothiphuc81299/coffeeShop-server/account/service"
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
)

// InitAdminServices ...
func InitAdminServices(d *model.CommonDAO) *model.AdminService {
	return &model.AdminService{
		Drink:    drinkservice.NewDrinkAdminService(d),
		Category: categoryservice.NewCategoryAdminService(d),
		User:     userservice.NewUserAdminService(d),
		Account:  accountservice.NewAccountAdminService(d),
		Role:     roleservice.NewRoleAdminService(d),
		Order:    orderservice.NewOrderService(d),
		Feedback: feedbackservice.NewFeedbackAdminService(d),
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
		Account:  accountDAO.NewAccountDAO(db),
		User:     userDAO.NewUserDAO(db),
		Role:     roleDAO.NewRoleDAO(db),
		Order:    orderDAO.NewOrderDAO(db),
		Feedback: feedbackDAO.NewFeedbackDAO(db),
	}
}
