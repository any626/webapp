package controller

import (
	// "fmt"
	"github.com/any626/webapp/service"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("views/*"))

type Controller struct {
	Service *service.Service
}

type Page struct {
	Title string
	Error string
	Auth  bool
}

func NewController(service *service.Service) *Controller {
	controller := &Controller{Service: service}

	return controller
}

func renderTemplate(w http.ResponseWriter, tmpl string, page Page) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
