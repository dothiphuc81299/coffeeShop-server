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
				bson.M{
					"_id":  "category_view",
					"name": "View",
				},
				bson.M{
					"_id":  "category_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "category_delete",
					"name": "Delete",
				},
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
				bson.M{
					"_id":  "drink_view",
					"name": "View",
				},
				bson.M{
					"_id":  "drink_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "drink_delete",
					"name": "Delete",
				},
			},
		},

		// 3. feedback
		bson.M{
			"_id":  "feedback",
			"name": "feedback",
			"permissions": []interface{}{
				bson.M{
					"_id":     "feedback_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				bson.M{
					"_id":  "feedback_view",
					"name": "View",
				},
				bson.M{
					"_id":  "feedback_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "feedback_delete",
					"name": "Delete",
				},
			},
		},
		// 4. shift
		bson.M{
			"_id":  "shift",
			"name": "shift",
			"permissions": []interface{}{
				bson.M{
					"_id":     "shift_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				bson.M{
					"_id":  "shift_view",
					"name": "View",
				},
				bson.M{
					"_id":  "shift_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "shift_delete",
					"name": "Delete",
				},
			},
		},
		// 4. salary
		bson.M{
			"_id":  "salary",
			"name": "salary",
			"permissions": []interface{}{
				bson.M{
					"_id":     "salary_admin",
					"name":    "Admin",
					"isAdmin": true,
				},
				bson.M{
					"_id":  "salary_view",
					"name": "View",
				},
				bson.M{
					"_id":  "salary_edit",
					"name": "Edit",
				},
				bson.M{
					"_id":  "salary_delete",
					"name": "Delete",
				},
			},
		},
	}
)
