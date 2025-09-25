package repository

import (
	"assignment/openapi"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	ID          string     `gorm:"primarykey"`
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

func CreateMovie(movie *openapi.Movie) error {
	m := Movie{
		Title:       movie.Title,
		Genre:       movie.Genre,
		ReleaseDate: movie.ReleaseDate.Time,
		Distributor: movie.Distributor,
		Budget:      movie.Budget,
		MpaRating:   movie.MpaRating,
	}

	m.ID = fmt.Sprintf("m_%s", uuid.New().String())
	movie.Id = m.ID

	return DB.Create(&m).Error
}

func GetMovies(params openapi.GetMoviesParams) ([]Movie, string, error) {
	query := DB.Model(&Movie{})

	if params.Q != nil {
		query = query.Where("title LIKE ?", "%"+*params.Q+"%")
	}

	if params.Year != nil {
		query = query.Where("YEAR(release_date) = ?", *params.Year)
	}

	if params.Genre != nil {
		query = query.Where("LOWER(genre) = LOWER(?)", *params.Genre)
	}

	if params.Distributor != nil {
		query = query.Where("LOWER(distributor) = LOWER(?)", *params.Distributor)
	}

	if params.Budget != nil {
		query = query.Where("budget <= ?", *params.Budget)
	}

	if params.MpaRating != nil {
		query = query.Where("mpa_rating = ?", *params.MpaRating)
	}

	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}

	if params.Cursor != nil {
		query = query.Where("id > ?", *params.Cursor)
	}

	var movies []Movie
	if err := query.Find(&movies).Error; err != nil {
		return nil, "", err
	}

	nextCursor := ""
	if len(movies) > 0 {
		nextCursor = fmt.Sprintf("%s", movies[len(movies)-1].ID)
	}

	return movies, nextCursor, nil
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
