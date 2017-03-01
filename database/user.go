package database

import (
    // "fmt"
    // "log"
    "database/sql"
)

type User struct {
    ID         int64
    Email      string
    Password   string
    FirstName sql.NullString
    LastName  sql.NullString
}

func (db *DB) CreateUser(user *User) error {
    _, err := db.DB.Exec(`INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`,
        user.Email, user.Password, user.FirstName, user.LastName)
    if err != nil {
        return err
    }
    return nil
}