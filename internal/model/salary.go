package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SalaryRaw struct {
	User        primitive.ObjectID   `bson:"user"`
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

type SalaryAdminService interface {
	GetDetail(ctx context.Context, salary SalaryBody, staff StaffRaw) SalaryResponse
	GetMonth(cond bson.M, month string)
}