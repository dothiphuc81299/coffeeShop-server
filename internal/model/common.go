package model

import (
	"strings"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AppID custom ObjectID
type AppID = primitive.ObjectID

// NewAppID ...
func NewAppID() AppID {
	return primitive.NewObjectID()
}

// CommonQuery ...
type CommonQuery struct {
	Page      int64
	Limit     int64
	Sort      bson.D
	Keyword   string
	Status    string
	Active    string
	Brand     AppID
	Partner   AppID
	Warehouse AppID
}

// AssignKeyword ...
func (q *CommonQuery) AssignKeyword(cond *bson.M) {
	if q.Keyword != "" {
		q.Keyword = format.NonAccentVietnamese(q.Keyword)
		(*cond)["searchString"] = format.SearchString(strings.Trim(q.Keyword, " "))
	}
}

// AssignActive ...
func (q *CommonQuery) AssignActive(cond *bson.M) {
	if q.Active != "" {
		if q.Active == "true" {
			(*cond)["active"] = true
		}
		if q.Active == "false" {
			(*cond)["active"] = false
		}
	}
}

// ResponseAdminListData ...
type ResponseAdminListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage"`
}

// ResponseAdminData ...
type ResponseAdminData struct {
	Data interface{} `json:"data"`
}
