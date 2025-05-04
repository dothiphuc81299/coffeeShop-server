package rest

import "github.com/dothiphuc81299/coffeeShop-server/pkg/order/category"

type Server struct {
	Dependences *Dependences
}

type Dependences struct {
	CategorySrv category.Service
}
