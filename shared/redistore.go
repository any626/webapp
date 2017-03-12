package shared

import (
    "net/http"
    "github.com/garyburd/redigo/redis"
    redistore "gopkg.in/boj/redistore.v1"
    "github.com/gorilla/sessions"
    "log"
)

type RediStore struct {
    RStore     *redistore.RediStore
    sessionKey string
}

func (rStore *RediStore) Get(r *http.Request) (*sessions.Session, error) {
    return rStore.RStore.Get(r, rStore.sessionKey)
}


func NewRediStoreWithPool(redisPool *redis.Pool, authKey []byte) *RediStore {
    rStore, err := redistore.NewRediStoreWithPool(redisPool, authKey)
    if err != nil {
        log.Fatalln(err)
    }
    return &RediStore{RStore: rStore, sessionKey: "session"}
}