package database

import (
    "testing"
    "database/sql"
    "gopkg.in/DATA-DOG/go-sqlmock.v1"
    "errors"
)

func TestCreateUser(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Errorf("%s", err)
        return
    }
    defer db.Close()

    database := &DB{DB: db}

    user := &User{Email: "test@test.com", Password: "password", FirstName: sql.NullString{String: "first", Valid: true}, LastName: sql.NullString{String: "last", Valid: true}}

    mock.ExpectExec("INSERT INTO users").WithArgs("test@test.com", "password", "first", "last").WillReturnResult(sqlmock.NewResult(1, 2))

    err = database.CreateUser(user)
    if err != nil {
        t.Errorf("%s", err)
    }

    mock.ExpectExec("INSERT INTO users").WithArgs("test@test.com", "password", "first", "last").WillReturnError(errors.New("duplicate"))
    err = database.CreateUser(user)
    if err.Error() != "duplicate" {
        t.Errorf("%s", err)
    }

    // we make sure that all expectations were met
    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("There were unfulfilled expections: %s", err)
    }
}



// // CreateUser creates the user in the database.
// func (db *DB) CreateUser(user *User) error {
//     _, err := db.DB.Exec(`INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`,
//         user.Email, user.Password, user.FirstName, user.LastName)
//     if err != nil {
//         return err
//     }
//     return nil
// }