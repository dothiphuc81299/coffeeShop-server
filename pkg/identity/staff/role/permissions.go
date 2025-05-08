package role

const (
	PermissionView   = "view"
	PermissionUpdate = "update"
	PermissionDelete = "delete"
	PermissionCreate = "create"
	PermissionAdmin  = "admin"
)

const (
	ResourceUser      = "user"
	ResourceDrink     = "drink"
	ResourceCategory  = "category"
	ResourceOrder     = "order"
	ResourceEvent     = "event"
	ResourceStaff     = "staff"
	ResourceRole      = "role"
	ResourceStatistic = "statistic"
)

var Permissions = []string{
	// Category
	"category_admin",
	"category_view",
	"category_update",
	"category_delete",
	"category_create",

	// Drinks
	"drink_admin",
	"drink_view",
	"drink_update",
	"drink_delete",
	"drink_create",

	// Orders
	"order_admin",
	"order_view",
	"order_update",

	// Events
	"event_admin",
	"event_view",
	"event_update",
	"event_create",
	"event_delete",

	// Users
	"user_admin",
	"user_view",
	"user_update",

	// Statistic
	"statistic_admin",

	// Staff
	"staff_view",

	// Role
	"role_view",
}
