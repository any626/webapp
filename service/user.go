package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
	"golang.org/x/crypto/bcrypt"
	// "errors"
)

func (s *Service) CreateUser(email, password string) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &database.User{Email:email, Password: string(hashPassword)}
	return s.DB.CreateUser(user)
}