package entity

type Movie struct {
	Budget           int      `json:"budget,omitempty"`
	Genres           []string `json:"genres"`
	Homepage         string   `json:"homepage,omitempty"`
	ID               int      `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	Overview         string   `json:"overview"`
	ReleaseDate      string   `json:"release_date"`
	Runtime          int      `json:"runtime,omitempty"`
	Title            string   `json:"title"`
	Similar          []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"similar"`
	Writers []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"writers"`
	Cast []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"cast"`
}
