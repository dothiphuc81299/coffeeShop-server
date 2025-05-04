package rest

import (
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"
	"github.com/dothiphuc81299/coffeeShop-server/pkg/order/drink"
)

type Server struct {
	Dependences *Dependences
}

type Dependences struct {
	CategorySrv category.Service
	DrinkSrv    drink.Service
}
