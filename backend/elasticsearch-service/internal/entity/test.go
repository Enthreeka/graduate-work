package entity

type Test struct {
	ID int `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Address     string `json:"address"`
	MAIL        string `json:"mail"`
	Name        string `json:"name"`
}
