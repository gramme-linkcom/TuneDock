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

func nullStringToPtr(v sql.NullString) *string {
	if !v.Valid {
		return nil
	}

	return &v.String
}
func GetTracks(database *sql.DB) ([]model.Track, error) {
	rows, err := database.Query(`
		SELECT
			id,
			album_id,
			title,
			disc_number,
			track_number,
			duration_ms,
			file_path,
			file_size,
			cover_id,
			codec,
			bit_depth,
			sample_rate,
			bitrate,
			is_lossless,
			is_hires
		FROM tracks
		ORDER BY album_id, disc_number, track_number, title
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := []model.Track{}

	for rows.Next() {
		var t model.Track

		var albumID sql.NullString
		var trackNumber sql.NullInt64
		var durationMS sql.NullInt64
		var fileSize sql.NullInt64
		var coverID sql.NullString
		var codec sql.NullString
		var bitDepth sql.NullInt64
		var sampleRate sql.NullInt64
		var bitrate sql.NullInt64
		var isLossless int
		var isHiRes int

		err := rows.Scan(
			&t.ID,
			&albumID,
			&t.Title,
			&t.DiscNumber,
			&trackNumber,
			&durationMS,
			&t.FilePath,
			&fileSize,
			&coverID,
			&codec,
			&bitDepth,
			&sampleRate,
			&bitrate,
			&isLossless,
			&isHiRes,
		)
		if err != nil {
			return nil, err
		}

		t.AlbumID = nullStringToPtr(albumID)
		t.TrackNumber = nullIntToPtr(trackNumber)
		t.DurationMS = nullIntToPtr(durationMS)
		t.BitDepth = nullIntToPtr(bitDepth)
		t.SampleRate = nullIntToPtr(sampleRate)
		t.Bitrate = nullIntToPtr(bitrate)

		if fileSize.Valid {
			t.FileSize = fileSize.Int64
		}

		if coverID.Valid {
			t.CoverID = coverID.String
		}

		if codec.Valid {
			t.Codec = codec.String
		}

		t.IsLossless = isLossless == 1
		t.IsHiRes = isHiRes == 1

		tracks = append(tracks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tracks, nil
}

func GetTrackByID(database *sql.DB ,id uuid.UUID) (*model.Track, error){
	row := database.QueryRow(`
		SELECT
			id,
			title,
			disc_number,
			track_number,
			duration_ms,
			file_path,
			file_size,
			cover_id,
			codec,
			bit_depth,
			sample_rate,
			bitrate,
			is_lossless,
			is_hires
		FROM tracks
		WHERE id = ?
	`, id)

	var t model.Track

	var albumID sql.NullString
	var coverID sql.NullString

	var trackNumber sql.NullInt64
	var durationMS sql.NullInt64
	var fileSize sql.NullInt64

	var codec sql.NullString
	var bitDepth sql.NullInt64
	var sampleRate sql.NullInt64
	var bitrate sql.NullInt64

	var isLossless int
	var isHiRes int

	err := row.Scan(
		&t.ID,
		//&albumID,
		&t.Title,
		&t.DiscNumber,
		&trackNumber,
		&durationMS,
		&t.FilePath,
		&fileSize,
		&coverID,
		&codec,
		&bitDepth,
		&sampleRate,
		&bitrate,
		&isLossless,
		&isHiRes,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	t.AlbumID = nullStringToPtr(albumID)
	t.TrackNumber = nullIntToPtr(trackNumber)
	t.DurationMS = nullIntToPtr(durationMS)
	if coverID.Valid {
		t.CoverID = coverID.String
	}

	if fileSize.Valid {
		t.FileSize = fileSize.Int64
	}

	if codec.Valid {
		t.Codec = codec.String
	}

	t.BitDepth = nullIntToPtr(bitDepth)
	t.SampleRate = nullIntToPtr(sampleRate)
	t.Bitrate = nullIntToPtr(bitrate)

	t.IsLossless = isLossless == 1
	t.IsHiRes = isHiRes == 1

	return &t, nil
}

func UpsertTrack(database *sql.DB, track *model.Track) error {
	_, err := database.Exec(`
		INSERT INTO tracks (
			id,
			title,
			album_id,
			disc_number,
			track_number,
			duration_ms,
			file_path,
			file_size,
			codec,
			bit_depth,
			sample_rate,
			bitrate,
			is_lossless,
			is_hires,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
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
		track.AlbumID,
		track.DiscNumber,
		track.TrackNumber,
		track.DurationMS,
		track.FilePath,
		track.FileSize,
		track.Codec,
		track.BitDepth,
		track.SampleRate,
		track.Bitrate,
		boolToInt(track.IsLossless),
		boolToInt(track.IsHiRes),
	)

	return err
}
