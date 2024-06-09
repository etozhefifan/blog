package main

import (
	"net/http"
	"html/template"
	"log"
	"time"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
)

  // Контейнер для компилированного шаблона
var tmpl *template.Template

func handler(w http.ResponseWriter, r *http.Request) {
    // Парсим шаблон
    tmpl, err := template.ParseFiles("/bin/template.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Данные для передачи в шаблон как карта
    data := map[string]string{
        "Title":   "Пример рендеринга шаблонов в Go",
        "Header":  "This is a header",
	"TimeStamp": time.Now().Format(time.DateTime),
        "Message": "Это пример рендеринга HTML шаблона в Go.",
    }

    // Рендерим шаблон
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func initializeDatabase() *gorm.DB {
    host := os.Getenv("POSTGRES_HOST")
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")
    port := os.Getenv("POSTGRES_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    log.Println("Успешно подключено к базе данных!")
    return db
}

func main() {
    http.HandleFunc("/", handler)
    db := initializeDatabase()
    log.Printf(db.Migrator().CurrentDatabase())
    log.Println("Сервер запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
