package database

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	APIKey    string `json:"apikey"`
	IsManager bool   `json:"ismanager"`
}
