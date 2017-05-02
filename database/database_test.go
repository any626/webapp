package database

import (
    "testing"
)

func TestCreateDNS(t *testing.T) {
    config := &Config{Host:"localhost", Port: 999, Database: "webapp", Username: "test", Password: "password", SslMode: "disable"}
    dns := createDNS(config)
    if dns != "user=test dbname=webapp port=999 host=localhost sslmode=disable password=password" {
        t.Fatal("Incorrect dns")
    }
    
}