package util

import (
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAppIDFromHex ...
func GetAppIDFromHex(s string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(s)
	return id
}

// ConvertObjectIDsToStrings ...
func ConvertObjectIDsToStrings(ids []primitive.ObjectID) []string {
	return funk.Map(ids, func(item primitive.ObjectID) string {
		return item.Hex()
	}).([]string)
}

// ConvertStringsToObjectIDs ...
func ConvertStringsToObjectIDs(strValues []string) []primitive.ObjectID {
	return funk.Map(strValues, func(item string) primitive.ObjectID {
		return GetAppIDFromHex(item)
	}).([]primitive.ObjectID)
}
