package models

type Sport struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	BetEnabled           bool   `json:"bet_enabled"`
	Enabled              bool   `json:"enabled"`
	CategoryID           int    `json:"category_id"`
	GoldMedalCountryID   int    `json:"gold_medal_country_id"`
	SilverMedalCountryID int    `json:"silver_medal_country_id"`
	BronzeMedalCountryID int    `json:"bronze_medal_country_id"`

	Category           SportCategory `gorm:"foreignkey:CategoryID;references:ID" json:"category"`
	GoldenMedalCountry Country       `gorm:"foreignkey:GoldMedalCountryID;references:ID" json:"golden_medal_country"`
	SilverMedalCountry Country       `gorm:"foreignkey:SilverMedalCountryID;references:ID" json:"silver_medal_country"`
	BronzeMedalCountry Country       `gorm:"foreignkey:BronzeMedalCountryID;references:ID" json:"bronze_medal_country"`
}
