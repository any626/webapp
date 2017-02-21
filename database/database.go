package database

import (
    "log"
    // "io/ioutil"
    // "encoding/json"
    "fmt"
    "database/sql"
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

func Connect(c *Config) *sql.DB {
    fmt.Println("Connecting to Database...")
    db, err := sql.Open("postgres", createDNS(c))
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println("Testing database connection")
    err = db.Ping()
    if err != nil {
        log.Fatalln(err)
    }
    
    fmt.Println("Connected to database successfully")
    return db
}

func createDNS(c *Config) string {
    return fmt.Sprintf("user=%s dbname=%s port=%d host=%s sslmode=%s password=%s", c.Username, c.Database, c.Port, c.Host, c.SslMode, c.Password)
}