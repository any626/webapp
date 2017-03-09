package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
	"golang.org/x/crypto/bcrypt"
	// "errors"
	// "time"
	// "github.com/garyburd/redigo/redis"
 //    "gopkg.in/boj/redistore.v1"
    "net/http"
    // "github.com/gorilla/sessions"
)

var SessionKey string = "session-key"

func (s *Service) CreateUser(user *database.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashPassword)
	return s.DB.CreateUser(user)
}

func (s *Service) SignIn(w http.ResponseWriter, r *http.Request) {
	email    := r.FormValue("email")
	password := r.FormValue("password")

	user := s.DB.GetUserByEmail(email)

	if user == nil {
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := s.RediStore.Get(r, SessionKey)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	session.Values["userId"] = user.ID
	session.Save(r, w)

	http.Redirect(w, r, "/home", 302)
}
