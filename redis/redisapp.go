package redisapp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/logrusorgru/aurora"
)

var (
	store *redis.Client
)

// Connect connect to redis server
func Connect(uri, pwd string) {
	ctx := context.Background()
	store = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: pwd,
		DB:       0, // use default DB
	})

	// Test
	_, err := store.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Cannot connect to redis", uri, err)
	}
	
	fmt.Println(aurora.Green("*** CONNECTED TO REDIS: " + uri))
}

// GetClient ...
func GetClient() *redis.Client {
	return store
}

// SetKeyValue ...
func SetKeyValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	storeByte, _ := json.Marshal(value)
	r := store.Set(ctx, key, storeByte, expiration)
	return r.Err()
}

// GetValueByKey ...
func GetValueByKey(key string) string {
	ctx := context.Background()
	value, _ := store.Get(ctx, key).Result()
	return value
}

// DelKey ...
func DelKey(key string) error {
	ctx := context.Background()
	return store.Del(ctx, key).Err()
}

// DelAllKeyByPattern ...
func DelAllKeyByPattern(pattern string) error {
	ctx := context.Background()
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = store.Scan(ctx, cursor, pattern, 100000).Result()
		if err != nil {
			return err
		}

		// If found any keys, delete
		if len(keys) > 0 {
			if err := store.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}
	return nil
}
