package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
	"github.com/any626/webapp/shared"
	"github.com/garyburd/redigo/redis"
	// "github.com/gorilla/sessions"
)

// Service is used to hold functionality that may be commonly shared.
type Service struct {
	DB        *database.DB
	RedisPool *redis.Pool
	RediStore *shared.RediStore
}

// NewService returns a Service type.
func NewService(db *database.DB, redisPool *redis.Pool, rediStore *shared.RediStore) *Service {
	service := &Service{DB: db, RedisPool: redisPool, RediStore: rediStore}

	return service
}
