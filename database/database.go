package database

import (
	"log"
	// "io/ioutil"
	// "encoding/json"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	SslMode  string `json:"sslmode"`
}

type DB struct {
	DB *sql.DB
}

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

func (db *DB) Disconnect() error {
	return db.DB.Close()
}

func createDNS(c *Config) string {
	return fmt.Sprintf("user=%s dbname=%s port=%d host=%s sslmode=%s password=%s", c.Username, c.Database, c.Port, c.Host, c.SslMode, c.Password)
}
