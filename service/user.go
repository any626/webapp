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
	// "fmt"
	// "github.com/gorilla/sessions"
)

func (s *Service) CreateUser(user *database.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashPassword)
	return s.DB.CreateUser(user)
}

func (s *Service) SignIn(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := s.DB.GetUserByEmail(email)

	if user == nil {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	session, err := s.RediStore.Get(r)
	if err != nil {
		http.Redirect(w, r, "/sign-in", http.StatusFound)
		return
	}

	session.Values["userId"] = user.ID
	err = session.Save(r, w)

	http.Redirect(w, r, "/home", http.StatusFound)
}

func (s *Service) IsAuth(w http.ResponseWriter, r *http.Request) bool {
	session, err := s.RediStore.Get(r)
	if err != nil {
		return false
	}

	_, ok := session.Values["userId"]
	if !ok {
		return false
	}

	return true
}
