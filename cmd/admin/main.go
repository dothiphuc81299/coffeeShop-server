package main

import (
	"github.com/dothiphuc81299/coffeeShop-server/initialize"
	"github.com/dothiphuc81299/coffeeShop-server/internal/config"
	"github.com/dothiphuc81299/coffeeShop-server/internal/locale"
	"github.com/dothiphuc81299/coffeeShop-server/internal/server"
)

func init() {
	locale.LoadProperties()
	config.Init()
}

func main() {
	_, commonDAO := initialize.ConnectDB(config.GetEnv().Mongo)
	service := initialize.InitAdminServices(commonDAO)

	e := server.StartAdmin(service, commonDAO)

	e.Logger.Fatal(e.Start(config.GetEnv().AdminPort))
}
