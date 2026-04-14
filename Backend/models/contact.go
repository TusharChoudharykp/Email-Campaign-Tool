package models

type Contact struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Segment   string `json:"segment"`
	CreatedAt string `json:"created_at"`
}