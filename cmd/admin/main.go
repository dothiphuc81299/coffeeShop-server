package main

import (
	"context"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/initialize"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/server"
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
	total := d.User.CountByCondition(ctx, bson.M{"isRoot": true})
	if total <= 0 {
		now := time.Now()
		// Init user root
		doc := model.UserRaw{
			ID:           primitive.NewObjectID(),
			Username:     "Root",
			SearchString: format.NonAccentVietnamese("Root"),
			Phone:        "0702654453",
			Active:       true,
			Avatar:       model.FileDefaultPhoto(),
			CreatedAt:    now,
			UpdatedAt:    now,
			IsRoot:       true,
		}

		d.User.InsertOne(ctx, doc)

		// Init account root
		accountDoc := model.AccountRaw{
			ID: model.NewAppID(),

			Active:      true,
			User:        doc.ID,
			Permissions: make([]string, 0),
			//		IsRoot:      true,
			CreatedAt: now,
			UpdatedAt: now,
		}
		d.Account.InsertOne(ctx, accountDoc)
	}
}
