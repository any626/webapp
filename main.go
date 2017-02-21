package main

import (
	"fmt"
	"log"
	"os"
	// "github.com/any626/webapp/routes"
	"github.com/any626/webapp/database"
	// "net/http"
	"encoding/json"
	"io/ioutil"
)

var environments []string = []string{"local", "staging", "production"}

var ENV string = os.Getenv("ENV")

func main() {

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

	fmt.Println(ENV)

	fmt.Println("=== Loading Configs ===")
	b, err := ioutil.ReadFile("./configs/"+ENV+"/config.json")
    if err != nil {
        log.Fatalln(err)
    }

    configs := configs{}
    err = json.Unmarshal(b, &configs)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Printf("%+v\n", configs)

	// load configs
	// database.Connect()
	// http.Handle("/", routes.Routes)
	// http.ListenAndService(":8080", nil)
}

type configs struct {
	database.Config `json:"Database"`
}