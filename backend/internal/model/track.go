package model

type Track struct {
	ID          string  `json:"id"`
	AlbumID     *string `json:"album_id,omitempty"`

	Title       string `json:"title"`
	DiscNumber  int    `json:"disc_number"`
	TrackNumber *int   `json:"track_number,omitempty"`
	DurationMS  *int   `json:"duration_ms,omitempty"`

	FilePath string `json:"-"`
	FileSize int64  `json:"file_size"`

	Codec      string `json:"codec"`
	BitDepth   *int   `json:"bit_depth,omitempty"`
	SampleRate *int   `json:"sample_rate,omitempty"`
	Bitrate    *int   `json:"bitrate,omitempty"`

	IsLossless bool `json:"is_lossless"`
	IsHiRes    bool `json:"is_hires"`

	Album  *Album   `json:"album,omitempty"`
	Artists []Artist `json:"artists,omitempty"`
}
