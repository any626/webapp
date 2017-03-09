package controller

import (
    // "fmt"
    "html/template"
    "github.com/any626/webapp/service"
    "net/http"
)

var SessionKey string = "session-key"

var templates = template.Must(template.ParseGlob("views/*"))

type Controller struct {
    Service *service.Service
}

type tmplData struct {
    Title string
    Error string
}

func NewController(service *service.Service) *Controller {
    controller := &Controller{Service: service}

    return controller
}

func renderTemplate(w http.ResponseWriter, tmpl string, data tmplData) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}