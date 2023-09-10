package models

type Sport struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	BetEnabled bool   `json:"bet_enabled"`
	Enabled    bool   `json:"enabled"`

	Category           SportCategory `json:"category"`
	GoldenMedalCountry Country       `json:"golden_medal_country"`
	SilverMedalCountry Country       `json:"silver_medal_country"`
	BronzeMedalCountry Country       `json:"bronze_medal_country"`
}
