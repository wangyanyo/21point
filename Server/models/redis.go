package models

import "github.com/redis/go-redis/v9"

var Cache *redis.Client

const (
	RankListCacheKey = "ranklist"
)

func InitRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
