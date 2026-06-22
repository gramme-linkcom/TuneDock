package api

import (
	"net/http"
	"os"
	"path/filepath"
	"github.com/google/uuid"
)

func TrackCoverHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		trackIdStr := query.Get("id")
		if trackIdStr == "" {
			http.Error(w, "track id is required", http.StatusBadRequest)
			return
		}

		_, err := uuid.Parse(trackIdStr)
		if err != nil {
			http.Error(w, "invalid track id", http.StatusBadRequest)
			return
		}

		coverPath := filepath.Join(
			"data", 
			"cache", 
			"cover", 
			trackIdStr+
			".webp")

		cover, err := os.Open(coverPath)
		if err != nil {
			http.Error(w, "cover not found", http.StatusNotFound)
			return
		}
		defer cover.Close()
		

		w.Header().Set("Content-Type", "image/webp;")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		http.ServeFile(w, r, coverPath)
	}
}
