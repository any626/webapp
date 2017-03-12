package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
    "github.com/garyburd/redigo/redis"
    "github.com/any626/webapp/shared"
    // "github.com/gorilla/sessions"
)

type Service struct {
    DB *database.DB
    RedisPool *redis.Pool
    RediStore *shared.RediStore
}

func NewService(db *database.DB, redisPool *redis.Pool, rediStore *shared.RediStore) *Service {
	service := &Service{DB: db, RedisPool: redisPool, RediStore: rediStore}

	return service
}