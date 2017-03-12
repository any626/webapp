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
    "fmt"
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

	session, err := s.RediStore.Get(r)
	if err != nil {
	    http.Error(w, err.Error(), http.StatusInternalServerError)
	    return
	}

	session.Values["userId"] = user.ID
	err = session.Save(r, w)
	fmt.Println(err)

	http.Redirect(w, r, "/home", 302)
}

func (s *Service) IsAuth(w http.ResponseWriter, r *http.Request) bool {
	session, err := s.RediStore.Get(r)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return false
    }

    _, ok := session.Values["userId"]
    if !ok {
        w.WriteHeader(http.StatusUnauthorized)
        return false
    }

    return true
}
