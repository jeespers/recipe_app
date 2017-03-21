package main

import (
	"html/template"
	"net/http"
	"os"
)

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

	http.ListenAndServe(":"+port, nil)
}
