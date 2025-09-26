package handler

import (
	"assignment/config"
	"assignment/internal/client"
	"assignment/internal/repository"
	"assignment/openapi"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"golang.org/x/sync/errgroup"
)

type MovieHandler struct {
	cli openapi.ClientWithResponsesInterface
}

var _ openapi.ServerInterface = &MovieHandler{}

func NewMovieHandler() *MovieHandler {
	cli, _ := openapi.NewClientWithResponses(config.Conf.BoxOfficeUrl, openapi.WithHTTPClient(client.NewBoxOfficeClient()))
	return &MovieHandler{cli: cli}
}

func (h *MovieHandler) GetMovies(c *gin.Context, params openapi.GetMoviesParams) {

	movies, nextCursor, err := repository.GetMovies(params)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	data := openapi.MoviePage{Items: make([]openapi.Movie, len(movies))}
	if nextCursor != "" {
		data.NextCursor = &nextCursor
	}

	g, _ := errgroup.WithContext(context.Background())

	for i := range movies {
		i := i
		data.Items[i] = openapi.Movie{
			Title:       movies[i].Title,
			Genre:       movies[i].Genre,
			ReleaseDate: openapi_types.Date{Time: movies[i].ReleaseDate},
			Distributor: movies[i].Distributor,
			Budget:      movies[i].Budget,
			MpaRating:   movies[i].MpaRating,
			Id:          movies[i].ID,
		}
		g.Go(func() error {
			response, err := h.cli.GetMovieBoxOfficeWithResponse(c, &openapi.GetMovieBoxOfficeParams{Title: "线些西周"})
			if err != nil {
				return err
			}
			if err = json.Unmarshal(response.Body, &data.Items[i].BoxOffice); err != nil {
				return err
			}
			return nil
		})
	}

	if err = g.Wait(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

func (h *MovieHandler) PostMovies(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "Bearer "+config.Conf.AuthToken {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	m := openapi.Movie{}

	if err := c.ShouldBindBodyWithJSON(&m); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	if m.Title == "" {
		c.JSON(422, gin.H{"error": "missing title"})
		return
	}

	if err := repository.CreateMovie(&m); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, m)
}

func (h *MovieHandler) GetMoviesTitleRating(c *gin.Context, title string) {
	stat, err := repository.GetAverageRating(title)
	if err != nil {
		return
	}

	if stat.Count == 0 {
		c.JSON(404, gin.H{"error": "No ratings found for this movie"})
		return
	}

	c.JSON(200, stat)
}

var validRating = map[float32]struct{}{
	0.5: {},
	1.0: {},
	1.5: {},
	2.0: {},
	2.5: {},
	3.0: {},
	3.5: {},
	4.0: {},
	4.5: {},
	5.0: {},
}

func (h *MovieHandler) PostMoviesTitleRatings(c *gin.Context, title string) {
	raterId := c.GetHeader("X-Rater-Id")

	if raterId == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	r := repository.Rating{}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		return
	}

	movie, err := repository.GetMovieByTitle(title)
	if err != nil {
		c.JSON(404, nil)
		return
	}

	if movie.ID == "" {
		c.JSON(404, gin.H{"error": "Movie not found"})
		return
	}

	if _, ok := validRating[r.Rating]; !ok {
		c.JSON(422, nil)
	}

	r.MovieTitle = title
	r.RaterId = raterId
	rowAffected, err := repository.UpsertRating(&r)
	if err != nil {
		c.JSON(400, nil)
		return
	}

	if rowAffected == 1 {
		c.JSON(201, r)
	} else {
		c.JSON(200, r)
	}

}
