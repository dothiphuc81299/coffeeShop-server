package model

// CommonDAO ....
type CommonDAO struct {
	Drink    DrinkDAO
	Category CategoryDAO
}

// AdminService ....
type AdminService struct {
	Drink    DrinkAdminService
	Category CategoryAdminService
}
