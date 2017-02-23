package main

import (
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/any626/webapp/controllers"
)

func routes(c *controllers.Controllers) *negroni.Negroni {
	r := mux.NewRouter()

	r.HandleFunc("/", c.HomeController.Home).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}