package main

import (
	"fmt"
	"log"
	"os"
	"github.com/any626/webapp/routes"
	"net/http"
	// "json"
)

var environments []string = []string{"local", "staging", "production"}

func main() {

	var env string = os.Getenv("ENV")

	isValidEnv := false
	for _, v := range environments {
		if v == env {
			isValidEnv = true
			break
		}
	}

	if !isValidEnv {
		log.Fatalln("Unknown Environment Variable. Shutting Down")
	}

	fmt.Println("=== Loading Configs ===")
	// load configs

	http.Handle("/", routes.Routes)
	http.ListenAndService(":8080", nil)

}