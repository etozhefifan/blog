package main

import (
	"net/http"
	"html/template"
	"log"
	"time"
//	"os"
//	"gorm.io/gorm"
//	"gorm.io/driver/postgres"
//	"fmt"
)

var tmpl *template.Template
var timestamp string

func handler(w http.ResponseWriter, r *http.Request) { 
    tmpl, err := template.ParseFiles("./index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := map[string]string{
        "Title":   "fifan.in",
        "Header":  "Hi there, traveler!",
	"Timestamp": timestamp,
	"Message": "Here's how you can contant me:",
	"Discord": "https://discordapp.com/users/188324670074781696",
    }

    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

//func initializeDatabase() *gorm.DB {
//    host := os.Getenv("POSTGRES_HOST")
//    user := os.Getenv("POSTGRES_USER")
//    password := os.Getenv("POSTGRES_PASSWORD")
//    dbname := os.Getenv("POSTGRES_DB")
//    port := os.Getenv("POSTGRES_PORT")
//    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
//    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//    if err != nil {
//        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
//    }

//    log.Println("Успешно подключено к базе данных!")
//    return db
//}

func main() {
    http.HandleFunc("/", handler)
    timestamp = time.Now().Format(time.DateTime)
//    db := initializeDatabase()
//    log.Printf(db.Migrator().CurrentDatabase())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
