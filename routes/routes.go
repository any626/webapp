package routes

import (
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/any626/webapp/controllers"
)

func Routes(c *controllers.Controllers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", c.HomeController.Home).Methods("GET")

	return r
}