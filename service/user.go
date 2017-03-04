package service

import (
	"fmt"
	"github.com/any626/webapp/database"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func (s *Service) CreateUser(user *database.User) error {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashPassword)
	return s.DB.CreateUser(user)
}

func (s *Service) createJWT(user *database.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.Auth.Key))
	return tokenString, err
}

func (s *Service) ParseJwt(t string) error {
	token, err := jwt.Parse(t, func (token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.Auth.Key), nil
	})

	if err != nil {
		fmt.Println(err)
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	    fmt.Println(claims)
	} else {
		return errors.New("Invalid token or claims")
	}

	return nil
}
