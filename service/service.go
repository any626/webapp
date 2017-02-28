package service

import (
	// "fmt"
	"github.com/any626/webapp/database"
)

type Service struct {
    DB *database.DB
}

func NewService(db *database.DB) *Service {
	service := &Service{DB: db}

	return service
}