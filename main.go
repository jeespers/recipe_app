package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Page struct for handling html page content
type Page struct {
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	page := &Page{Title: title}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%T", db)
	http.ListenAndServe(":"+port, nil)
}
