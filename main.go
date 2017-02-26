package main

import (
	"fmt"
	"log"
	"os"
	// "github.com/any626/webapp/routes"
	"github.com/any626/webapp/controllers"
	"github.com/any626/webapp/database"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var environments []string = []string{"local", "staging", "production"}

var ENV string = os.Getenv("ENV")

func main() {

	checkEnv()

	config := loadConfig()

	db := database.Connect(&config.Database)
	defer db.Disconnect()
	
	handlers := controllers.NewControllers(&db)

	// handlers.HomeController.Test()

	middlewaredRouter := routes(handlers)

	http.Handle("/", middlewaredRouter)
	http.ListenAndServe(":8080", nil)
}

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

func loadConfig() Config {
	fmt.Println("Loading configs...")
	b, err := ioutil.ReadFile("./configs/"+ENV+"/config.json")
    if err != nil {
        log.Fatalln(err)
    }

    config := Config{}
    err = json.Unmarshal(b, &config)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Loaded Configs")
    return config
}

type Config struct {
	Database database.Config `json:"database"`
}