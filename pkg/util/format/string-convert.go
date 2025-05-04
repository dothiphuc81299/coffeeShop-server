package format

import (
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const charset = "0123456789"

// StringToInt ...
func StringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// StringToFloat64 ...
func StringToFloat64(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return n
}

// StringToInt64 ...
func StringToInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// IntToString ...
func IntToString(i int) string {
	s := strconv.Itoa(i)
	return s
}

// SearchString ...
func SearchString(keyword string) bson.M {
	return bson.M{
		"$regex": bsonx.Regex(NonAccentVietnamese(keyword), "i"),
	}
}

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GenerateCodeString ...
func GenerateCodeString(length int) string {
	return stringWithCharset(length, charset)
}
