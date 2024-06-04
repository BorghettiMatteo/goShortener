package models

import (
	"github.com/redis/go-redis/v9"
)

var Database *redis.Client

func CreateDatabase() {
	Database = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
}
