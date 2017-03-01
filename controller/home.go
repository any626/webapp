package controller

import (
	"fmt"
	"net/http"
	"net/mail"
    // "html/template"
    "github.com/any626/webapp/database"
)

func (c *Controller) GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func (c *Controller) GetRegister(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "register", tmplData{Title: "Register"})
}

func (c *Controller) PostRegister(w http.ResponseWriter, r *http.Request) {
    email           := r.FormValue("email")
    password        := r.FormValue("password")
    confirmPassword := r.FormValue("confirm-password")

    _, err := mail.ParseAddress(email)

    if err != nil {
    	tData := tmplData{Title: "Register", Error: err.Error()}
    	renderTemplate(w, "register", tData)
    	return
    }
    if password != confirmPassword {
    	tData := tmplData{Title: "Register", Error: "Passwords do not match"}
    	renderTemplate(w, "register", tData)
    	return
    }

    user := &database.User{Email: email, Password: password}
    err = c.Service.CreateUser(user)

    if err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
    }

    fmt.Fprintf(w, "Registered")
}