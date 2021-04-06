package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/initialize"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/server"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	locale.LoadProperties()
	config.Init()
}

func main() {
	_, commonDAO := initialize.ConnectDB(config.GetEnv().Mongo)
	service := initialize.InitAdminServices(commonDAO)

	e := server.StartAdmin(service, commonDAO)
	// Init account admin root
	initAccountAdminRoot(commonDAO)

	e.Logger.Fatal(e.Start(config.GetEnv().AdminPort))
}

func initAccountAdminRoot(d *model.CommonDAO) {
	ctx := context.Background()
	// Check role root
	total := d.Staff.CountByCondition(ctx, bson.M{"isRoot": true})
	if total <= 0 {
		now := time.Now()
		// Init Account root
		doc := model.StaffRaw{
			ID:          primitive.NewObjectID(),
			Username: "admin",
			Password: "123456",
			Phone:       "0702654453",
			Active:      true,
			Role:        model.AppID{},
			Avatar:      model.FileDefaultPhoto(),
			CreatedAt:   now,
			UpdatedAt:   now,
			IsRoot:      true,
			Permissions: make([]string, 0),
		}

		d.Staff.InsertOne(ctx, doc)
		fmt.Println(aurora.Green("*** Init account admin root: "))
	}
}
