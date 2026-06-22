package api

import (
	// "encoding/json"
	"encoding/json"
	// "fmt"
	"net/http"
	"tunedock/internal/repository"
	"database/sql"

	"github.com/google/uuid"
)

func AlbumHandler(database *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		albumIdStr := query.Get("id")
		if albumIdStr == "" {
			http.Error(w, "track id is required", http.StatusBadRequest)
			return
		}

		albumId, err := uuid.Parse(albumIdStr)
		if err != nil {
			http.Error(w, "invalid album id", http.StatusBadRequest)
			return
		}
		
		album, err := repository.GetAlbumByID(database, albumId)
		if err != nil {
			http.Error(w, "failed to get album: "+err.Error(), http.StatusInternalServerError)
			return
		}

		Res := album

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if err := json.NewEncoder(w).Encode(Res); err != nil {
			http.Error(w, "Failed to encode json", http.StatusInternalServerError)
			return
		}
	}
}
