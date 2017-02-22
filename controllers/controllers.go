package controllers

import (
    // "fmt"
    "database/sql"
)

type Controllers struct {
    HomeController HomeController
    DB *sql.DB
}

func NewControllers(db *sql.DB) *Controllers {
    controllers := &Controllers{}

    return controllers
}