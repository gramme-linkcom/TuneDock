package api

import (
	"encoding/json"
	// "fmt"
	"database/sql"
	"net/http"
	"tunedock/internal/repository"

	"github.com/google/uuid"
)

func TrackHandler(database *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		trackIdStr := query.Get("id")
		if trackIdStr == "" {
			http.Error(w, "track id is required", http.StatusBadRequest)
			return
		}

		trackId, err := uuid.Parse(trackIdStr)
		if err != nil {
			http.Error(w, "invalid track id", http.StatusBadRequest)
			return
		}
		
		track, err := repository.GetTrackByID(database, trackId)
		if err != nil {
			http.Error(w, "failed to get track: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if track == nil {
			http.Error(w, "failed to get track: id is not found", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		if err := json.NewEncoder(w).Encode(track); err != nil {
			http.Error(w, "Failed to encode json", http.StatusInternalServerError)
			return
		}
	}
}

func TrackListsHandler(database *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		tracks, err := repository.GetTracks(database)
		if err != nil {
			http.Error(w, "failed to get tracks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(tracks); err != nil {
			http.Error(w, "failed to encode tracks", http.StatusInternalServerError)
			return
		}
	}
}
