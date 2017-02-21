package database

import (
    // "log"
    // "io/ioutil"
    // "encoding/json"
    // "fmt"
)

type Config struct {
    Host     string `json:"host"`
    Port     int    `json:"port"`
    Username string `json:"username"`
    Password string `json:"password"`
    Charset  string `json:"charset"`
    Schema   string `json:"schema"`
    SslMode  string `json:"sslmode"`

}

// func Connect() {
//     file, err := ioutil.ReadFile("/configs/"+ENV+"datbase.json")
//     if err != nil {
//         log.Fatalln(err)
//     }

//     config := Config{}

//     err = json.Unmarshal(file, &config)
//     if err != nil {
//         log.Fatalln(err)
//     }
//     fmt.Println(config)
// }