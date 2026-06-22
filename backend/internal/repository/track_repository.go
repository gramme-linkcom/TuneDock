package repository

import (
	"database/sql"
	"tunedock/internal/model"

	"github.com/google/uuid"
)

func nullIntToPtr(v sql.NullInt64) *int {
	if !v.Valid {
		return nil
	}

	i := int(v.Int64)
	return &i
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func GetTrackByID(database *sql.DB ,id uuid.UUID) (*model.Track, error){
	row := database.QueryRow(`
		SELECT
			id,
			title,
			codec,
			file_path,
			bit_depth,
			sample_rate,
			is_lossless,
			is_hires
		FROM tracks
		WHERE id = ?
	`, id)

	var t model.Track
	var bitDepth sql.NullInt64
	var sampleRate sql.NullInt64
	var isLossless int
	var isHiRes int

	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Codec,
		&t.FilePath,
		&bitDepth,
		&sampleRate,
		&isLossless,
		&isHiRes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	t.BitDepth = nullIntToPtr(bitDepth)
	t.SampleRate = nullIntToPtr(sampleRate)
	t.IsLossless = isLossless == 1
	t.IsHiRes = isHiRes == 1

	return &t, nil
}

func UpsertTrack(database *sql.DB, track *model.Track) error {
	_, err := database.Exec(`
		INSERT INTO tracks (
			id,
			title,
			file_path,
			file_size,
			codec,
			is_lossless,
			is_hires,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(file_path) DO UPDATE SET
			title = excluded.title,
			file_size = excluded.file_size,
			codec = excluded.codec,
			is_lossless = excluded.is_lossless,
			is_hires = excluded.is_hires,
			updated_at = CURRENT_TIMESTAMP
	`,
		uuid.NewString(),
		track.Title,
		track.FilePath,
		track.FileSize,
		track.Codec,
		boolToInt(track.IsLossless),
		boolToInt(track.IsHiRes),
	)

	return err
}
