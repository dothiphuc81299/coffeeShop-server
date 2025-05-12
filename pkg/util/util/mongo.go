package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObjectIDFromHex(s string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(s)
	return id
}
