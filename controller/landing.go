package controller

import (
	"fmt"
	"net/http"
	"net/mail"
    // "html/template"
    "github.com/any626/webapp/database"
)

func (c *Controller) GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index")
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

func (c *Controller) GetSignIn(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "signIn", tmplData{Title: "Sign In"})
}

func (c *Controller) PostSignIn(w http.ResponseWriter, r *http.Request) {
    c.Service.SignIn(w, r)
}

func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
    session, err := c.Service.RediStore.Get(r, SessionKey)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    session.Options.MaxAge = -1
    session.Save(r, w)
    http.Redirect(w, r, "/", 302)
}