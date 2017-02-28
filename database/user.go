package database

import (
    // "fmt"
    // "log"
)

type User struct {
    ID         int64
    Email      string
    Password   string
    FirstName string
    LastName  string
}

func (db *DB) CreateUser(user *User) error {
    _, err := db.DB.Exec(`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`, user.Email, user.Password)
    if err != nil {
        return err
    }
    return nil
}