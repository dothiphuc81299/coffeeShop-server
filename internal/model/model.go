package model

// CommonDAO ....
type CommonDAO struct {
	Drink    DrinkDAO
	Category CategoryDAO

	User UserDAO
	Role RoleDAO

	Order OrderDAO

	Event EventDAO

	Staff     StaffDAO
	StaffRole StaffRoleDAO
	Session   SessionDAO

	DrinkAnalytic    DrinkAnalyticDAO
	Question         QuestionDAO
	Group            GroupDAO
	Package          PackageDAO
	UserPackageGroup UserPackageGroupDAO
	PackageGroup     PackageGroupDAO
}

// AdminService ....
type AdminService struct {
	Drink    DrinkAdminService
	Category CategoryAdminService
	User     UserAdminService
	Role     RoleService

	Event EventAdminService

	StaffRole     StaffRoleAdminService
	Staff         StaffAdminService
	Order         OrderAdminService
	DrinkAnalytic DrinkAnalyticService
	Question      QuestionAdminService
	Group         GroupAdminService
	Package       PackageAdminService
	PackageGroup  PackageGroupAdminService
}

type AppService struct {
	User  UserAppService
	Order OrderAppService
	Staff StaffAppService
}

// ResponseAppListData ...
type ResponseAppListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage,omitempty"`
}
