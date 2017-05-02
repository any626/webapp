package main

import (
	"encoding/json"
	"fmt"
	"github.com/any626/webapp/controller"
	"github.com/any626/webapp/database"
	"github.com/any626/webapp/router"
	"github.com/any626/webapp/service"
	"github.com/any626/webapp/shared"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
	"github.com/gorilla/handlers"
)

// main is driver function
func main() {

	config := loadConfig()
	fmt.Println("Loaded configs.")

	db := database.Connect(&config.Database)
	defer db.Disconnect()
	fmt.Println("Connected to database.")

	redisPool := shared.GetRedisPool(&config.Redis)
	fmt.Println("Connected to redis.")

	rediStore := shared.NewRediStoreWithPool(redisPool, []byte(config.Auth.Key))
	defer rediStore.RStore.Close()

	s := service.NewService(db, redisPool, rediStore)

	c := controller.NewController(s)

	r := router.NewRouter(c, s)

	http.Handle("/", handlers.RecoveryHandler()(r))

	fmt.Println("Server running...")
	http.ListenAndServe(":8080", nil)
}

// loadConfig loads the configs
func loadConfig() Config {
	b, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	config := Config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

// Config is the master for all configs.
type Config struct {
	Env      string             `json:"env"`
	Database database.Config    `json:"database"`
	Auth     Authentication     `json:"auth"`
	Redis    shared.RedisConfig `json:"redis"`
}

// Authentication is the config to hold the auth secret key.
type Authentication struct {
	Key string
}
