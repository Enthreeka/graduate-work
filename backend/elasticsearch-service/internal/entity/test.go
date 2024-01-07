package entity

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	MAIL string `json:"mail"`

	Address     string `json:"address"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
