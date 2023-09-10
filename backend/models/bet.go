package models

type Bet struct {
	User                 User   `json:"user"`
	Sport                Sport  `json:"sport"`
	GoldMedalCountryID   string `json:"gold_medal_country_id"`
	SilverMedalCountryID string `json:"silver_medal_country_id"`
	BronzeMedalCountryID string `json:"bronze_medal_country_id"`
}
