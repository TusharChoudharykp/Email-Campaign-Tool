package models

type Campaign struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Subject   string `json:"subject"`
	Template  string `json:"template"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}