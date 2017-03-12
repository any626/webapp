package database

import (
	"log"
	// "io/ioutil"
	// "encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Config holds the config data for the database.
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	SslMode  string `json:"sslmode"`
}

// DB holds the functionality for the database.
type DB struct {
	DB *sql.DB
}

// Connect connects to the database. Is set to postgres.
func Connect(c *Config) *DB {
	db, err := sql.Open("postgres", createDNS(c))
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &DB{DB: db}
}

// Disconnects disconnects the database.
func (db *DB) Disconnect() error {
	return db.DB.Close()
}

// Created the connection string for the database.
func createDNS(c *Config) string {
	return fmt.Sprintf("user=%s dbname=%s port=%d host=%s sslmode=%s password=%s", c.Username, c.Database, c.Port, c.Host, c.SslMode, c.Password)
}
