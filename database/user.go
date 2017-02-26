package database

import (
    "fmt"
)

type User struct {
    ID         int64
    Email      string
    Password   string
    FirstName string
    LastName  string
}

func (db *DB) createUser(user User) (User, bool) {
    res, err := db.db.Exec(`INSERT INTO users (name, email, password, first_name, last_name) VALUES
        ($1, $2, $3, $4, $5) RETURNING id`, user.Email, user.Password, user.FirstName, user.LastName)

    success := true

    if err == nil {
        fmt.Println(err)
        success = false
    }

    user.ID, err = res.LastInsertId()

    if err == nil {
        fmt.Println(err)
        success = false
    }

    return user, success
}