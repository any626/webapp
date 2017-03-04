package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
)

type Service struct {
    DB *database.DB
    Auth *Auth
}

type Auth struct {
    Key string
}

func NewService(db *database.DB, auth *Auth) *Service {
	service := &Service{DB: db, Auth: auth}

	return service
}