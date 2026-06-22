package model

type Artist struct {
	ID       string  `json:"id"`
	Name     string `json:"name"`
	SortName string `json:"sort_name,omitempty"`
}
