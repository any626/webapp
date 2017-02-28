package main

import (
	// "net/http"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/any626/webapp/controller"
)

func routes(c *controller.Controller) *negroni.Negroni {
	r := mux.NewRouter()

	r.HandleFunc("/", c.GetHome).Methods("GET")
    r.HandleFunc("/register", c.GetRegister).Methods("GET")
    r.HandleFunc("/register", c.PostRegister).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}