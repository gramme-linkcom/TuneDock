package repository

import (
	"database/sql"
	"tunedock/internal/model"

	"github.com/google/uuid"
)

func GetAlbumByID(database *sql.DB, id uuid.UUID) (*model.Album, error){
	row := database.QueryRow(`
		SELECT
			id,
			title,
			year
		FROM albums
		WHERE id = ?
	`, id)

	var t model.Album
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Year,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &t, nil
}
