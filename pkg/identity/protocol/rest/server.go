package rest

import "github.com/dothiphuc81299/coffeeShop-server/pkg/identity/user"

type Server struct {
	Dependences *Dependences
}

type Dependences struct {
	UserSrv user.Service
}
