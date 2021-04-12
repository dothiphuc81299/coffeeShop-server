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
	Staff   primitive.ObjectID
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

func (q *CommonQuery) AssignStatus(cond *bson.M) {
	if q.Status != "" {
		(*cond)["status"] = q.Status
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

func (q *CommonQuery) AssignStaff(cond *bson.M) {
	(*cond)["staff"] = q.Staff
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
func (q *CommonQuery) AssignStartAtAndEndAt(cond *bson.M) {

	if !q.StartAt.IsZero() && !q.EndAt.IsZero() {
		q.StartAt = util.TimeStartOfDayInHCM(q.StartAt.AddDate(0, 0, 1))
		q.EndAt = util.TimeStartOfDayInHCM(q.EndAt)
		(*cond)["date"] = bson.M{
			"$gte": q.StartAt,
			"$lte": q.EndAt,
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
