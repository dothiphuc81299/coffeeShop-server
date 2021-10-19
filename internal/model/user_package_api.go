package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserPackageGroupRaw struct {
	ID             primitive.ObjectID `bson:"_id"`
	PackageID      primitive.ObjectID `bson:"packageId"`
	UserID         primitive.ObjectID `bson:"userId"`
	SubmissionTime time.Time          `bson:"submissionTime"`
	TotalPoint     float64            `bson:"totalPoint"`
	IsPass         bool               `bson:"isPass"`
}
