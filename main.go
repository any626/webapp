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
	"os"
	"github.com/gorilla/handlers"
)

// environments holds the valid environments.
var environments []string = []string{"local", "staging", "production"}

// ENV holds the current environment variables.
var ENV string = os.Getenv("ENV")

// main is driver function
func main() {

	checkEnv()

	config := loadConfig()

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
	http.ListenAndServe(":8080", nil)
}

// checkEnv determines if the environment is valid.
func checkEnv() {
	isValidEnv := false
	for _, v := range environments {
		if v == ENV {
			isValidEnv = true
			break
		}
	}

	if !isValidEnv {
		log.Fatalln("Unknown Environment Variable. Shutting Down")
	}

	fmt.Println("Environment: " + ENV)
}

// loadConfig loads the functions under the configs/{environment} folder.
func loadConfig() Config {
	b, err := ioutil.ReadFile("./configs/" + ENV + "/config.json")
	if err != nil {
		log.Fatalln(err)
	}

	config := Config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Loaded configs")
	return config
}

// Config is the master for all configs.
type Config struct {
	Database database.Config    `json:"database"`
	Auth     Authentication     `json:"auth"`
	Redis    shared.RedisConfig `json:"redis"`
}

// Authentication is the config to hold the auth secret key.
type Authentication struct {
	Key string
}
