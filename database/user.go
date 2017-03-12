package database

import (
	"fmt"
	// "log"
	"database/sql"
	"time"
)

// User is a user model.
type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName sql.NullString
	LastName  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUser creates the user in the database.
func (db *DB) CreateUser(user *User) error {
	_, err := db.DB.Exec(`INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}

// GetUserByEmail gets the user model by email.
func (db *DB) GetUserByEmail(email string) *User {
	user := &User{}
	err := db.DB.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}
