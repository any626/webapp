package controllers

import (
    // "fmt"
    "html/template"
    "github.com/any626/webapp/database"
    "net/http"
)


var templates = template.Must(template.ParseGlob("views/*"))

type Controllers struct {
    HomeController HomeController
    DB *database.DB
}

type tmplData struct {
    Title string
}

func NewControllers(db *database.DB) *Controllers {
    controllers := &Controllers{}

    return controllers
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}