package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TotalPoints int    `json:"total_points"`
}
