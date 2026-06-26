package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func SearchHandler(database *sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		q     := query.Get("q")
		if q == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]any{})
			return
		}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(q)
	}
}
