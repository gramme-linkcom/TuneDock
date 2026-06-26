package main

import (
	"html/template"
	"log"
	"net/http"
	"tunedock/internal/api"
	"tunedock/internal/db"
	"tunedock/internal/library"
	"tunedock/internal/system"
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
	system.Init()

	database, err := db.Open("./data/tunedock.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err := tmpl.Execute(w, map[string]any{
			"Tracks": tracks,
		})
		if err != nil {
			http.Error(w, "template error", http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /library/update", library.LibraryUpdateHandler(database))
	mux.HandleFunc("GET /api/albums/detail", api.AlbumHandler(database))
	
	mux.HandleFunc("GET /api/tracks", api.TrackListsHandler(database))
	mux.HandleFunc("GET /api/tracks/detail", api.TrackHandler(database))
	mux.HandleFunc("GET /api/tracks/data/cover", api.TrackCoverHandler(database))
	
	mux.HandleFunc("GET /api/stream", api.TrackFileServeHandler(database))

	mux.HandleFunc("GET /api/search", api.SearchHandler(database))
	
	mux.Handle("/static/",http.StripPrefix("/static/", http.FileServer(http.Dir("../frontend/static/"))),)

	log.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
