package config

import "time"

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
