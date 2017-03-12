package database

import (
    "testing"
    // "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// func TestConnect(*testing.T) {
//     type Config struct {
//         Host     string `json:"host"`
//         Port     int    `json:"port"`
//         Database string `json:"database"`
//         Username string `json:"username"`
//         Password string `json:"password"`
//         SslMode  string `json:"sslmode"`
//     }

//     Connect(config)
    
// }

func TestCreateDNS(t *testing.T) {
    config := &Config{Host:"localhost", Port: 999, Database: "webapp", Username: "test", Password: "password", SslMode: "disable"}
    dns := createDNS(config)
    if dns != "user=test dbname=webapp port=999 host=localhost sslmode=disable password=password" {
        t.Fail()
    }
    
}