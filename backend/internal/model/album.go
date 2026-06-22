package model

type Album struct {
	ID        string  `json:"id"`
	Title     string `json:"title"`
	Year      *int   `json:"year,omitempty"`
	CoverPath string `json:"-"`

	Artists []Artist `json:"artists,omitempty"`
}
