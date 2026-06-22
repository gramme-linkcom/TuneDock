package api

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"tunedock/internal/repository"

	"github.com/google/uuid"
)

func contentTypeFromAudioPath(path string) string {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".flac":
		return "audio/flac"
	case ".mp3":
		return "audio/mpeg"
	case ".m4a", ".mp4":
		return "audio/mp4"
	case ".wav":
		return "audio/wav"
	case ".ogg", ".oga":
		return "audio/ogg"
	case ".opus":
		return "audio/opus"
	default:
		return "application/octet-stream"
	}
}

func TrackFileServeHandler(database *sql.DB) http.HandlerFunc{
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

		audioFilePath := filepath.Join(trackData.FilePath)
		// audioFilePath := filepath.Join("data", "music", "test", "audio.m4a")
		audioFile, err := os.Open(audioFilePath)
		if err != nil {
			http.Error(w, "track not found", http.StatusNotFound)
			return
		}
		defer audioFile.Close()
		
		fileInfo, err := audioFile.Stat()
		if err != nil {
			http.Error(w, "failed to read audio file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", contentTypeFromAudioPath(audioFilePath))
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Cache-Control", "private, max-age=3600")
		http.ServeContent(
			w,
			r,
			filepath.Base(audioFilePath),
			fileInfo.ModTime(),
			audioFile,
		)
	}
}
