package main

import (
	"html/template"
	"log"
	"net/http"
)

type Track struct {
	ID     int
	Title  string
	Artist string
}

var tracks = []Track{
	{ID: 1, Title: "Sample Track", Artist: "Unknown Artist"},
}

var tmpl = template.Must(template.ParseFiles("../frontend/index.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, map[string]any{
			"Tracks": tracks,
		})
		if err != nil {
			http.Error(w, "template error", http.StatusInternalServerError)
		}
	})

	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
