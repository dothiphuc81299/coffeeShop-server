package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryRaw struct {
	Staff       primitive.ObjectID   `bson:"staff"`
	Shift       []primitive.ObjectID `bson:"shift"`
	TotalShift  float64              `bson:"totalShift"`
	TotalSalary float64              `bson:"totalSalary"`
	Coefficient float64              `bson:"coefficient"`
	Allowance   float64              `bson:"allowance"`
	StartAt     time.Time            `bson:"startAt"`
	EndAt       time.Time            `bson:"endAt"`
}

type SalaryDAO interface {
}

type SalaryAppService interface {
	GetDetail(ctx context.Context, salary SalaryBody, staff StaffRaw) SalaryResponse
	//GetMonth(cond bson.M, month string)
}

type SalaryAdminService interface {
	GetList(ctx context.Context) []SalaryResponse
}
