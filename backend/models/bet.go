package models

type Bet struct {
	UserID               int `gorm:"primarykey" json:"user_id"`
	SportID              int `gorm:"primarykey" json:"sport_id"`
	GoldMedalCountryID   int `json:"gold_medal_country_id"`
	SilverMedalCountryID int `json:"silver_medal_country_id"`
	BronzeMedalCountryID int `json:"bronze_medal_country_id"`

	User               User    `gorm:"foreignkey:UserID;references:ID" json:"user"`
	Sport              Sport   `gorm:"foreignkey:SportID;references:ID" json:"sport"`
	GoldMedalCountry   Country `gorm:"foreignkey:GoldMedalCountryID;references:ID" json:"gold_medal_country"`
	SilverMedalCountry Country `gorm:"foreignkey:SilverMedalCountryID;references:ID" json:"silver_medal_country"`
	BronzeMedalCountry Country `gorm:"foreignkey:BronzeMedalCountryID;references:ID" json:"bronze_medal_country"`
}
