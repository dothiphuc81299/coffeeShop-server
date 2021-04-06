package model

// CommonDAO ....
type CommonDAO struct {
	Drink    DrinkDAO
	Category CategoryDAO

	User     UserDAO
	Role     RoleDAO
	Feedback FeedbackDAO
	Order    OrderDAO

	Event EventDAO

	Shift ShiftDAO

	Staff     StaffDAO
	StaffRole StaffRoleDAO
	Session   SessionDAO
}

// AdminService ....
type AdminService struct {
	Drink    DrinkAdminService
	Category CategoryAdminService
	User     UserAdminService
	Role     RoleService
	Feedback FeedbackAdminService
	Order    OrderAdminService

	Event EventAdminService

	Shift ShiftAdminService

	StaffRole StaffRoleAdminService
	Staff     StaffAdminService
}

// ResponseAppListData ...
type ResponseAppListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage,omitempty"`
}
