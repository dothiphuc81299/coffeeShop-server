package config

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Route const
const (
	AdminRoutePrefix = "" // change to "/admin" to add /admin prefix
)

// Header keys
const (
	HeaderOrigin         = "Origin"
	HeaderContentLength  = "Content-Length"
	HeaderContentType    = "Content-Type"
	HeaderAuthorization  = "Authorization"
	HeaderAcceptLanguage = "Accept-Language"
)

// Lang
const (
	LangVi = "vi"
	LangEn = "en"
)

// Time
const (
	TimezoneHCM              = "Asia/Ho_Chi_Minh"
	TimeDurationBetweenViews = time.Second * 30
)

// Regex
const (
	RegexPhoneNumber = `^(0|\+84|84)\d{9}$`
)

// Limit
const (
	Limit20 = 20
)

// ExtAllow
const (
	ExtAllow        = ".xlsx"
	FormatTime      = "2006_01_02_15_04_05_0700"
	PathUploadAdmin = "/uploads/"
)

const (
	ModelFieldUser     = "user"
	ModelFieldDrink    = "drink"
	ModelFieldCategory = "category"

	ModelFieldEvent = "event"
	ModelFieldOrder = "order"
	ModelFieldStaff = "staff"

	PermissionView   = "view"
	PermissionEdit   = "edit"
	PermissionDelete = "delete"
	PermissionCreate ="create"
	PermissionChangePassword ="changePassword"

	PermissionAdmin = "admin"
)

// Permissions ...
var (
	Permissions = []interface{}{
		// 1. Category
		bson.M{
			"_id":  "category",
			"name": "category",
			"permissions": []interface{}{
				bson.M{
					"_id":     "category_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				//bson.M{
				//	"_id":  "category_view",
				//	"name": "View",
				//},
				//bson.M{
				//	"_id":  "category_edit",
				//	"name": "Edit",
				//},
				//bson.M{
				//	"_id":  "category_delete",
				//	"name": "Delete",
				//},
			},
		},

		// 2. drinks
		bson.M{
			"_id":  "drinks",
			"name": "drinks",
			"permissions": []interface{}{
				bson.M{
					"_id":     "drink_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				//bson.M{
				//	"_id":  "drink_view",
				//	"name": "View",
				//},
				//bson.M{
				//	"_id":  "drink_edit",
				//	"name": "Edit",
				//},
				//bson.M{
				//	"_id":  "drink_delete",
				//	"name": "Delete",
				//},
			},
		},

		// 5.User
		bson.M{
			"_id":  "users",
			"name": "Users",
			"permissions": []interface{}{
				bson.M{
					"_id":     "user_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				bson.M{
					"_id":  "user_view",
					"name": "View",
				},
				bson.M{
					"_id":  "user_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "user_delete",
					"name": "Delete",
				},
			},
		},

		//6. Event
		bson.M{
			"_id":  "events",
			"name": "Events",
			"permissions": []interface{}{
				bson.M{
					"_id":     "event_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				//bson.M{
				//	"_id":  "event_view",
				//	"name": "View",
				//},
				//bson.M{
				//	"_id":  "event_edit",
				//	"name": "Edit",
				//},
			},
		},

		// 7.Order
		bson.M{
			"_id":  "orders",
			"name": "Orders",
			"permissions": []interface{}{
				bson.M{
					"_id":     "order_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				//bson.M{
				//	"_id":  "order_view",
				//	"name": "View",
				//},
				//bson.M{
				//	"_id":  "order_edit",
				//	"name": "Edit",
				//},
			},
		},

		// 7.Staff
		bson.M{
			"_id":  "staff",
			"name": "staff",
			"permissions": []interface{}{
				bson.M{
					"_id":     "staff_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				bson.M{
					"_id":  "staff_view",
					"name": "View",
				},
				bson.M{
					"_id":  "staff_edit",
					"name": "Edit",
				},
			},
		},
	}
)

