package model

type Album struct {
	ID        string  `json:"id"`
	Title     string `json:"title"`
	Year      *int   `json:"year,omitempty"`
	CoverID   string `json:"cover_id,omitempty"`

	Artists []Artist `json:"artists,omitempty"`
}
