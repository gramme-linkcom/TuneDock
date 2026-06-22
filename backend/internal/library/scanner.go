package library

import (
	"database/sql"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"

	"tunedock/internal/repository"
	"tunedock/internal/system"
)

var coverGenerationRunning atomic.Bool

func isAudioFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".mp3", ".flac", ".m4a", ".mp4", ".alac", ".ogg", ".oga", ".opus", ".wav", ".aiff", ".aif":
		return true
	default:
		return false
	}
}

func LibraryUpdateHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		config, err := system.GetConfig()
		if err != nil {
			http.Error(w, "Could not get config", http.StatusInternalServerError)
			return
		}

		for _, libraryData := range config.LibraryDatas {
			if !libraryData.Enabled {
				continue
			}

			scanPath(database, libraryData.Path)
		}
	}
}

func scanPath(database *sql.DB, root string) {
	root = filepath.Clean(root)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if !isAudioFile(path) {
			return nil
		}

		track, err := ReadAudioMetadata(path)
		if err != nil {
			return nil
		}

		if err := repository.UpsertTrack(database, track); err != nil {
			return nil
		}

		return nil
	})

	if err != nil {
		return
	}
}


func startGenerateCover(database *sql.DB) bool{
	if !coverGenerationRunning.CompareAndSwap(false, true) {
		return false
	}

	go func() {
		defer coverGenerationRunning.Store(false)
		generateCovers(database)
	}()
	return true
}

func generateCovers(database *sql.DB) {
	
}
