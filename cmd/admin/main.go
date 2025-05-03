package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/initialize"
	//"github.com/dothiphuc81299/coffeeShop-server/internal/config"

	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/model"
	"github.com/dothiphuc81299/coffeeShop-server/internal/server"
	"github.com/logrusorgru/aurora"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	locale.LoadProperties()
	//env :=ENV["REDIS_"]
	//uri := os.Getenv("REDIS_PROVIDER")
	// uri = URI.parse(ENV["REDIS_PROVIDER"])
	///fmt.Println("uri", uri)
	//uri, pass := ParseRedistogoUrl()
	//redisapp.Connect(uri, "")
	//	config.Init()
}

func ParseRedistogoUrl() (string, string) {
	redisUrl := os.Getenv("REDIS_URL")
	redisInfo, _ := url.Parse(redisUrl)
	server := redisInfo.Host
	password := ""
	if redisInfo.User != nil {
		password, _ = redisInfo.User.Password()
	}
	return server, password
}

func main() {
	_, commonDAO := initialize.ConnectDB()
	//	_, commonDAO := initialize.ConnectDB(config.GetEnv().Mongo)
	service := initialize.InitAdminServices(commonDAO)
	serviceApp := initialize.InitAppService(commonDAO)

	e := server.StartAdmin(service, serviceApp, commonDAO)
	// Init account admin root
	initAccountAdminRoot(commonDAO)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
	//	e.Logger.Fatal(e.Start(":" + "8082"))
}

//onst PORT = process.env.PORT || 4000

const avtDefault = "https://banner2.cleanpng.com/20180402/ojw/kisspng-united-states-avatar-organization-information-user-avatar-5ac20804a62b58.8673620215226654766806.jpg"

func initAccountAdminRoot(d *model.CommonDAO) {
	ctx := context.Background()
	// Check role root
	total := d.Staff.CountByCondition(ctx, bson.M{"isRoot": true})
	if total <= 0 {
		now := time.Now()
		// Init Account root
		doc := model.Staff{
			ID:          primitive.NewObjectID(),
			Username:    "admin",
			Password:    "123456",
			Phone:       "0702654453",
			Active:      true,
			Role:        model.AppID{},
			Avatar:      avtDefault,
			CreatedAt:   now,
			UpdatedAt:   now,
			IsRoot:      true,
			Permissions: make([]string, 0),
		}

		d.Staff.InsertOne(ctx, doc)
		fmt.Println(aurora.Green("*** Init account admin root: "))
	}
}
