package api

import (
	"database/sql"
	"net/http"
	//"os"

	//"path/filepath"
	"tunedock/internal/repository"

	"github.com/google/uuid"
	"go.senan.xyz/taglib"
)

func TrackCoverHandler(database *sql.DB) http.HandlerFunc{
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

		trackData, err := repository.GetTrackByID(database, trackId)
		if err != nil {
			http.Error(w, "failed to get track: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if trackData == nil {
			http.Error(w, "failed to get track: id is not found", http.StatusBadRequest)
			return
		}

		// coverPath := filepath.Join(
		// 	"data", 
		// 	"cache", 
		// 	"cover", 
		// 	trackIdStr+
		// 	".webp")

		coverRaw, err := taglib.ReadImage(trackData.FilePath)
		if err != nil {
			http.Error(w, "cover not found", http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "image/webp;")
		w.Header().Set("Cache-Control", "public, max-age=86400")
		if _, err := w.Write(coverRaw); err != nil {
			return
		}
		
	}
}
