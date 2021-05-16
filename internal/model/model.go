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

	Salary        SalaryDAO
	DrinkAnalytic DrinkAnalyticDAO
}

// AdminService ....
type AdminService struct {
	Drink    DrinkAdminService
	Category CategoryAdminService
	User     UserAdminService
	Role     RoleService

	Event EventAdminService

	Shift ShiftAdminService

	StaffRole StaffRoleAdminService
	Staff     StaffAdminService
	Order     OrderAdminService
	Feedback  FeedbackAdminService

	Salary        SalaryAdminService
	DrinkAnalytic DrinkAnalyticService
}

type AppService struct {
	User     UserAppService
	Feedback FeedbackAppService
	Order    OrderAppService
	Staff    StaffAppService
	Salary   SalaryAppService
}

// ResponseAppListData ...
type ResponseAppListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage,omitempty"`
}
