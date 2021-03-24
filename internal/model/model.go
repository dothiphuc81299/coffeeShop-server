package model

// CommonDAO ....
type CommonDAO struct {
	Drink    DrinkDAO
	Category CategoryDAO
	Account  AccountDAO
	User     UserDAO
	Role     RoleDAO
	Feedback FeedbackDAO
	Order    OrderDAO
}

// AdminService ....
type AdminService struct {
	Drink    DrinkAdminService
	Category CategoryAdminService
	Account  AccountAdminService
	User     UserAdminService
	Role     RoleService
	Feedback FeedbackAdminService
	Order    OrderAdminService
}

// ResponseAppListData ...
type ResponseAppListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage,omitempty"`
}
