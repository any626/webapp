package shared

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

// RedisConfig holds the config data for redis
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database int    `json:"database"`
}

// GetRedisPool returns a redis pool provided the RedisConfig
func GetRedisPool(config *RedisConfig) *redis.Pool {

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	return &redis.Pool{
		MaxIdle:   80, // max idle connections
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if err != nil {
				log.Fatalln(err.Error())
			}

			// Sets the database
			c.Do("SELECT", config.Database)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
