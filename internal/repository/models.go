package repository

import (
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Title       string     `gorm:"size:255;not null"`
	Genre       string     `gorm:"size:255;not null"`
	ReleaseDate time.Time  `gorm:"not null"`
	Distributor *string    `gorm:"size:255"`
	Budget      *int64     `gorm:"type:bigint"`
	MpaRating   *string    `gorm:"size:10"`
	BoxOffice   *BoxOffice `gorm:"type:json"`
}

func (Movie) TableName() string {
	return "movies"
}

type BoxOffice struct {
	Revenue     Revenue   `json:"revenue"`
	Currency    string    `json:"currency"`
	Source      string    `json:"source"`
	LastUpdated time.Time `json:"lastUpdated"`
}
type Revenue struct {
	Worldwide        int64  `json:"worldwide"`
	OpeningWeekendUs *int64 `json:"openingWeekendUsa,omitempty"`
}

type Rating struct {
	gorm.Model
	MovieTitle string  `gorm:"size:255;not null"`
	RaterId    string  `gorm:"size:255;not null"`
	Rating     float32 `gorm:"type:decimal(3,1);not null"`
}

func (Rating) TableName() string {
	return "ratings"
}
