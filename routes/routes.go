package routes

import (
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/any626/webapp/controllers"
)

func Routes() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomeController.getHome).Methods("GET")

	return r
}