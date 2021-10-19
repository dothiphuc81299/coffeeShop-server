package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// LoyaltyProgramUserStatusBSON ...
	LoyaltyProgramUserStatusBSON struct {
		ID             primitive.ObjectID `bson:"_id"`
		UserID         primitive.ObjectID `bson:"userID"`
		CurrentExpense float64            `bson:"currentExpense"`
		UpdatedAt      time.Time          `bson:"updatedAt"`
		CreatedAt      time.Time          `bson:"createdAt"`
	}
)
