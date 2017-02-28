package controllers

import (
	"net/http"
	"fmt"
    // "html/template"
)

type HomeController struct {
}

func (c *HomeController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home");
}

func (c *HomeController) GetRegister(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "register", tmplData{Title: "Register"})
}

func (c *HomeController) PostRegister(w http.ResponseWriter, r *http.Request) {
    // TODO
}