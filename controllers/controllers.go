package controllers

import (
    // "fmt"
    "github.com/any626/webapp/database"
)

type Controllers struct {
    HomeController HomeController
    DB *database.DB
}

func NewControllers(db *database.DB) *Controllers {
    controllers := &Controllers{}

    return controllers
}