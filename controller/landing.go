package controller

import (
	"fmt"
	"net/http"
	"net/mail"
	// "html/template"
	"github.com/any626/webapp/database"
)

// GetIndex is the GET handler for the index page.
func (c *Controller) GetIndex(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Welcome"}
	if c.Service.IsAuth(w, r) {
		page.Auth = true
	}
	renderTemplate(w, "index", page)
}

// GetRegister is the GET handler for the registration page.
func (c *Controller) GetRegister(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "register", Page{Title: "Register"})
}

// PostRegister is the POST handler for the registration page.
func (c *Controller) PostRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm-password")

	_, err := mail.ParseAddress(email)

	if err != nil {
		tData := Page{Title: "Register", Error: err.Error()}
		renderTemplate(w, "register", tData)
		return
	}
	if password != confirmPassword {
		tData := Page{Title: "Register", Error: "Passwords do not match"}
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

// GetSignIn is the GET handler for the sign in page.
func (c *Controller) GetSignIn(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "sign-in", Page{Title: "Sign In"})
}

// PostSignIn is the Post handler for the sign in page.
func (c *Controller) PostSignIn(w http.ResponseWriter, r *http.Request) {
	c.Service.SignIn(w, r)
}

// Logout is the handler for the logout functionality
func (c *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := c.Service.RediStore.Get(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1 // expire session
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)
}
