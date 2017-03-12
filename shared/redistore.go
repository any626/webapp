package shared

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/sessions"
	redistore "gopkg.in/boj/redistore.v1"
	"log"
	"net/http"
)

// SessionKey is the key to retrieve the user session
const SessionKey string = "session"

// RediStore holds the redistore and helping functions for ease of use.
type RediStore struct {
	RStore     *redistore.RediStore
}

// Get is a wrapper for redistore.get using a constant session key.
func (rStore *RediStore) Get(r *http.Request) (*sessions.Session, error) {
	// uses constant session key
	return rStore.RStore.Get(r, SessionKey)
}

// NewRediStoreWithPool returns a new RediStore struct provided a redis pool and an auth key.
func NewRediStoreWithPool(redisPool *redis.Pool, authKey []byte) *RediStore {
	rStore, err := redistore.NewRediStoreWithPool(redisPool, authKey)
	if err != nil {
		log.Fatalln(err)
	}
	return &RediStore{RStore: rStore}
}
