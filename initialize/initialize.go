package initialize

import (
	"context"
	"fmt"
	"time"

	categoryDAO "github.com/dothiphuc81299/coffeeShop-server/category/dao"
	categoryservice "github.com/dothiphuc81299/coffeeShop-server/category/service"
	drinkDAO "github.com/dothiphuc81299/coffeeShop-server/drink/dao"
	drinkservice "github.com/dothiphuc81299/coffeeShop-server/drink/service"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"

	gameDAO "github.com/dothiphuc81299/coffeeShop-server/game/dao"

	orderDAO "github.com/dothiphuc81299/coffeeShop-server/order/dao"
	orderservice "github.com/dothiphuc81299/coffeeShop-server/order/service"
	userDAO "github.com/dothiphuc81299/coffeeShop-server/user/dao"
	userservice "github.com/dothiphuc81299/coffeeShop-server/user/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	eventDAO "github.com/dothiphuc81299/coffeeShop-server/event/dao"

	eventservice "github.com/dothiphuc81299/coffeeShop-server/event/service"

	gameservice "github.com/dothiphuc81299/coffeeShop-server/game/service"
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

		Event: eventservice.NewEventAdminService(d),

		StaffRole:    staffroleservice.NewStaffRoleAdminService(d),
		Staff:        staffservice.NewStaffAdminService(d),
		Order:        orderservice.NewOrderAdminService(d),
		Question:     gameservice.NewQuestionAdminService(d),
		Group:        gameservice.NewGroupAdminService(d),
		Package:      gameservice.NewPackageAdminService(d),
		PackageGroup: gameservice.NewPackageGroupAdminService(d),
	}
}

func InitAppService(d *model.CommonDAO) *model.AppService {
	return &model.AppService{
		User:  userservice.NewUserAppService(d),
		Order: orderservice.NewOrderAppService(d),
		Staff: staffservice.NewStaffAppService(d),
	}
}

// ConnectDB ..
func ConnectDB() (*mongo.Database, *model.CommonDAO) {
	//client, err := mongo.NewClient(options.Client().ApplyURI(dbCfg.URI))
	//	dbCfg config.MongoCfg
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://coffeeShop:coffeeShop@cluster0.puhkn.mongodb.net/Cluster0?retryWrites=true&w=majority"))

	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	//db := client.Database(dbCfg.Name)
	db := client.Database("CoffeeShop")

	return db, &model.CommonDAO{
		Drink:    drinkDAO.NewDrinkDAO(db),
		Category: categoryDAO.NewCategoryDAO(db),

		User:    userDAO.NewUserDAO(db),
		CodeDAO: userDAO.NewCodedRegisterDAO(db),
		Order:   orderDAO.NewOrderDAO(db),

		Event:            eventDAO.NewEventDAO(db),
		Staff:            staffDAO.NewStaffDAO(db),
		StaffRole:        staffroleDAO.NewStaffRoleDAO(db),
		Session:          staffDAO.NewSessionDAO(db),
		Package:          gameDAO.NewPackageDAO(db),
		Question:         gameDAO.NewQuestionDAO(db),
		PackageGroup:     gameDAO.NewPackageGroupDAO(db),
		UserPackageGroup: gameDAO.NewUserPackageGroupDAO(db),
		Group:            gameDAO.NewGroupDAO(db),
	}
}
