package main

import (
	"fmt"
	"log"
	"os"
	"github.com/any626/webapp/controller"
	"github.com/any626/webapp/database"
	"github.com/any626/webapp/service"
	"github.com/any626/webapp/router"
	"encoding/json"
	"io/ioutil"
	"net/http"
	// "html/template"
)

var environments []string = []string{"local", "staging", "production"}

var ENV string = os.Getenv("ENV")

func main() {

	checkEnv()

	config := loadConfig()

	db := database.Connect(&config.Database)
	defer db.Disconnect()
	
	s := service.NewService(db, &config.Auth)

	c := controller.NewController(s)

	r := router.NewRouter(c, s)

	http.Handle("/", r)
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
	Auth service.Auth `json:"auth"`
}