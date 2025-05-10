package format

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func SearchString(keyword string) bson.M {
	return bson.M{
		"$regex": bsonx.Regex(NonAccentVietnamese(keyword), "i"),
	}
}
