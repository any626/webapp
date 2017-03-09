package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
    "github.com/garyburd/redigo/redis"
    rStore "gopkg.in/boj/redistore.v1"
    // "github.com/gorilla/sessions"
)

type Service struct {
    DB *database.DB
    RedisPool *redis.Pool
    RediStore *rStore.RediStore
}

func NewService(db *database.DB, redisPool *redis.Pool, rediStore *rStore.RediStore) *Service {
	service := &Service{DB: db, RedisPool: redisPool, RediStore: rediStore}

	return service
}