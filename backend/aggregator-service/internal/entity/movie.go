package entity

import "time"

type Movie struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Overview         string    `json:"overview"`
	ReleaseDate      time.Time `json:"release_date"`
	Runtime          int       `json:"runtime"`
	Budget           float64   `json:"budget"`
	OriginalLanguage string    `json:"original_language"`
	TrailerYT        string    `json:"trailer_yt"`
}
