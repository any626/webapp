package controller

import (
	// "fmt"
	"github.com/any626/webapp/service"
	"html/template"
	"net/http"
)

// templates holds the cached templates
var templates = template.Must(template.ParseGlob("views/*"))

// Page is the data used for templates
type Page struct {
    Title string
    Error string
    Auth  bool
}

// Controller holds the http handlers
type Controller struct {
	Service *service.Service
}

// NewController returns a Controller
func NewController(service *service.Service) *Controller {
	controller := &Controller{Service: service}

	return controller
}

// renderTemplate renders the templates
func renderTemplate(w http.ResponseWriter, tmpl string, page Page) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
