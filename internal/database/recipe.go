package database

type Recipe struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Base        string   `json:"base"`
	Temp        string   `json:"temp"`
	Ingredients []string `json:"ingredients"`
}
