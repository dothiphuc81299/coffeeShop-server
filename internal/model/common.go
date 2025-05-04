package model

// import (
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/dothiphuc81299/coffeeShop-server/internal/format"
// 	"github.com/dothiphuc81299/coffeeShop-server/internal/util"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // AppID custom ObjectID
// type AppID = primitive.ObjectID

// // NewAppID ...
// func NewAppID() AppID {
// 	return primitive.NewObjectID()
// }

// // CommonQuery ...
// type CommonQuery struct {
// 	Page     int64
// 	Limit    int64
// 	Sort     bson.D
// 	Keyword  string
// 	Status   string
// 	Active   string
// 	Staff    primitive.ObjectID
// 	IsCheck  string
// 	StartAt  time.Time
// 	EndAt    time.Time
// 	Category string
// 	Month    string
// 	Username string
// }

// func (q *CommonQuery) GetFindOptions() *options.FindOptions {
// 	opts := options.Find()
// 	if q.Limit > 0 {
// 		opts.SetLimit(q.Limit).SetSkip(q.Limit * q.Page)
// 	}
// 	if len(q.Sort) > 0 {
// 		opts = opts.SetSort(q.Sort)
// 	}
// 	return opts
// }

// // AssignKeyword ...
// func (q *CommonQuery) AssignKeyword(cond *bson.M) {
// 	if q.Keyword != "" {
// 		q.Keyword = format.NonAccentVietnamese(q.Keyword)
// 		(*cond)["searchString"] = format.SearchString(strings.Trim(q.Keyword, " "))
// 	}
// }

// func (q *CommonQuery) AssignUsername(cond *bson.M) {
// 	if q.Username != "" {
// 		q.Username = format.NonAccentVietnamese(q.Username)
// 		(*cond)["username"] = format.SearchString(strings.Trim(q.Username, " "))
// 	}
// }

// func (q *CommonQuery) AssignStatus(cond *bson.M) {
// 	if q.Status != "" {
// 		(*cond)["status"] = q.Status
// 	}
// }

// // AssignActive ...
// func (q *CommonQuery) AssignActive(cond *bson.M) {
// 	// if q.Active != "" {
// 	// 	if q.Active == "true" {
// 	// 		(*cond)["active"] = true
// 	// 	}
// 	// 	if q.Active == "false" {
// 	// 		(*cond)["active"] = false
// 	// 	}
// 	// }

// 	if q.Active != "" && q.Active != "all" {
// 		b, _ := strconv.ParseBool(q.Active)
// 		(*cond)["active"] = b
// 	}
// }

// func (q *CommonQuery) AssignStaff(cond *bson.M) {
// 	(*cond)["staff"] = q.Staff
// }

// func (q *CommonQuery) AssignIsCheck(cond *bson.M) {
// 	if q.IsCheck != "" {
// 		if q.IsCheck == "true" {
// 			(*cond)["isCheck"] = true
// 		}
// 		if q.IsCheck == "false" {
// 			(*cond)["isCheck"] = false
// 		}
// 	}
// }

// // AssignStartAtAndEndAt ...
// func (q *CommonQuery) AssignStartAtAndEndAt(cond *bson.M) {
// 	if !q.StartAt.IsZero() && !q.EndAt.IsZero() {
// 		q.StartAt = util.TimeStartOfDayInHCM(q.StartAt.AddDate(0, 0, 1))
// 		q.EndAt = util.TimeStartOfDayInHCM(q.EndAt)
// 		(*cond)["date"] = bson.M{
// 			"$gte": q.StartAt,
// 			"$lte": q.EndAt,
// 		}
// 	}

// }

// func (q *CommonQuery) AssignStartAtAndEndAtForDrink(cond *bson.M) {
// 	if !q.StartAt.IsZero() && !q.EndAt.IsZero() {
// 		q.StartAt = util.TimeStartOfDayInHCM(q.StartAt.AddDate(0, 0, 1))
// 		q.EndAt = util.TimeStartOfDayInHCM(q.EndAt)
// 		(*cond)["createdAt"] = bson.M{
// 			"$gte": q.StartAt,
// 			"$lte": q.EndAt,
// 		}
// 	}
// }

// func (q *CommonQuery) AssignStartAtAndEndAtByStatistic(cond *bson.M) {
// 	if q.StartAt.IsZero() && !q.EndAt.IsZero() {
// 		q.StartAt = q.BeginningOfMonth(q.EndAt)
// 	}

// 	if q.EndAt.IsZero() && !q.StartAt.IsZero() {
// 		q.EndAt = q.EndOfMonth(q.StartAt)
// 	}

// 	if q.StartAt.IsZero() && q.EndAt.IsZero() {
// 		q.StartAt, q.EndAt = q.GetThisMonthNow()
// 	}

// 	// q.StartAt = util.TimeStartOfDayInHCM(q.StartAt)
// 	// q.EndAt = util.TimeStartOfDayInHCM(q.EndAt)
// 	(*cond)["createdAt"] = bson.M{
// 		"$gte": q.StartAt,
// 		"$lte": q.EndAt,
// 	}
// }

// func (q *CommonQuery) GetThisMonthNow() (time.Time, time.Time) {
// 	now := time.Now()
// 	currentYear, currentMonth, _ := now.Date()
// 	currentLocation := now.Location()

// 	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
// 	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

// 	return firstOfMonth, lastOfMonth
// }

// func (q *CommonQuery) BeginningOfMonth(date time.Time) time.Time {
// 	now := time.Now()
// 	dDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now.UTC().Location())
// 	firstDay := dDate.AddDate(0, 0, -date.Day()+1)
// 	return firstDay
// }

// func (q *CommonQuery) EndOfMonth(date time.Time) time.Time {
// 	now := time.Now()
// 	dDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, now.UTC().Location())
// 	lastDay := dDate.AddDate(0, 1, -dDate.Day()+1).Add(time.Nanosecond * -1)
// 	return lastDay
// }

// func (q *CommonQuery) AssignCategory(cond *bson.M) {
// 	if q.Category != "" {
// 		categoryID, _ := primitive.ObjectIDFromHex(q.Category)
// 		if !categoryID.IsZero() {
// 			(*cond)["category"] = categoryID
// 		}
// 	}
// }

// ResponseAdminListData ...
type ResponseAdminListData struct {
	Data         interface{} `json:"data"`
	Total        int64       `json:"total"`
	LimitPerPage int64       `json:"limitPerPage"`
}

// // ResponseAdminData ...
// type ResponseAdminData struct {
// 	Data interface{} `json:"data"`
// }

// // GetFindOptsUsingPage ...
// func (q *CommonQuery) GetFindOptsUsingPage() *options.FindOptions {
// 	opts := options.Find()
// 	if q.Limit > 0 {
// 		opts.SetLimit(q.Limit).SetSkip((q.Page) * q.Limit)
// 	}
// 	if q.Sort != nil {
// 		opts.SetSort(q.Sort)
// 	}
// 	return opts
// }

// func (q *CommonQuery) GetFindOptsUsingPageOne() *options.FindOptions {
// 	opts := options.Find()
// 	if q.Limit > 0 {
// 		opts.SetLimit(q.Limit).SetSkip((q.Page - 1) * q.Limit)
// 	}
// 	if q.Sort != nil {
// 		opts.SetSort(q.Sort)
// 	}
// 	return opts
// }

// // GetFindOptsUsingTimestamp ...
// func (q *CommonQuery) GetFindOptsUsingTimestamp() *options.FindOptions {
// 	opts := options.Find()
// 	if q.Limit > 0 {
// 		opts.SetLimit(q.Limit)
// 	}
// 	if q.Sort != nil {
// 		opts.SetSort(q.Sort)
// 	}
// 	return opts
// }

// // GetFindOptionsUsingSort ...
// func (q *CommonQuery) GetFindOptionsUsingSort() *options.FindOptions {
// 	return options.Find().SetSort(q.Sort)
// }
