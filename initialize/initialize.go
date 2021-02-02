package initialize

import (
	"context"
	"fmt"
	"time"

	drinkDAO "github.com/dothiphuc81299/coffeeShop-server/drink/dao"
	drinkservice "github.com/dothiphuc81299/coffeeShop-server/drink/service"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitAdminServices ...
func InitAdminServices(d *model.CommonDAO) *model.AdminService {
	return &model.AdminService{
		Drink: drinkservice.NewDrinkAdminService(d),
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
		Drink: drinkDAO.NewDrinkDAO(db),
	}
}
