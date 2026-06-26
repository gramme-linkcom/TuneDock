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
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &t, nil
}

func GetAlbumByName(database *sql.DB, name string) (*model.Album, error){
	row := database.QueryRow(`
		SELECT
			id,
			title,
			year
		FROM albums
		WHERE title = ?
	`, name)

	var t model.Album
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Year,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return &t, nil
}

func UpsertAlbum(database *sql.DB, album *model.Album) error {
	_, err := database.Exec(`
		INSERT INTO albums (
			id,
			title,
			year,
			cover_id,
			updated_at
		)
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
	`,
		album.ID,
		album.Title,
		album.Year,
		album.CoverID,
	)

	return err
}
