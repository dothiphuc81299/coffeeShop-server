package useraccount

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAccountRaw struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	LoginName    string             `bson:"login_name" json:"login_name"`
	Active       bool               `bson:"active" json:"active"`
	CurrentPoint float64            `bson:"current_point" json:"current_point"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
}

type CreateUserAccountCommand struct {
	UserID    primitive.ObjectID `json:"user_id"`
	LoginName string             `json:"login_name"`
	Active    bool               `json:"active"`
}

