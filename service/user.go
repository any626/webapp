package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
	"golang.org/x/crypto/bcrypt"
	// "errors"
	// "github.com/dgrijalva/jwt-go"
)

func (s *Service) CreateUser(user *database.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashPassword)
	return s.DB.CreateUser(user)
}