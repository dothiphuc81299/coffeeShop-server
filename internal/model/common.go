package model

import (
	"strings"
	"time"

	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
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
	Page    int64
	Limit   int64
	Sort    bson.D
	Keyword string
	Status  string
	Active  string
	User    primitive.ObjectID
	IsCheck string
	StartAt time.Time
	EndAt   time.Time
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

func (q *CommonQuery) AssignUser(cond *bson.M) {
	(*cond)["user"] = q.User
}

func (q *CommonQuery) AssignIsCheck(cond *bson.M) {
	if q.IsCheck != "" {
		if q.IsCheck == "true" {
			(*cond)["isCheck"] = true
		}
		if q.IsCheck == "false" {
			(*cond)["isCheck"] = false
		}
	}
}

// AssignStartAtAndEndAt ...
func (q CommonQuery) AssignStartAtAndEndAt(cond *bson.M) {
	check := time.Time{}
	if q.StartAt != check && q.EndAt != check {
		(*cond)["date"] = bson.M{
			"$gte": util.TimeStartOfDayInHCM(q.StartAt),
			"$lte": util.TimeEndOfDayHCM(q.EndAt),
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
