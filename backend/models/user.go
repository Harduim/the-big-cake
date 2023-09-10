package models

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TotalPoints int    `json:"total_points"`
	Enabled     bool   `json:"enabled"`
	OidSSO      string `json:"oid_sso"`
}
