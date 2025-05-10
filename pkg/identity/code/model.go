package code

import "go.mongodb.org/mongo-driver/bson/primitive"

type CodedRegisterRaw struct {
	Id    primitive.ObjectID `bson:"_id"`
	Email string             `bson:"email"`
	Code  string             `bson:"code"`
}
