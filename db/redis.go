package db

import (
	"github.com/go-redis/redis"
)

var rds *redis.Client

// Init initializes the redis by populating the client
func Init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// GetDB returns the redis client instance
func GetDB() *redis.Client {
	return rds
}
