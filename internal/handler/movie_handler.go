package handler

import (
	"assignment/config"
	"assignment/internal/client"
	"assignment/internal/repository"
	"assignment/openapi"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	openapi_types "github.com/oapi-codegen/runtime/types"
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

	for i := range movies {
		data.Items[i] = openapi.Movie{
			Title:       movies[i].Title,
			Genre:       movies[i].Genre,
			ReleaseDate: openapi_types.Date{Time: movies[i].ReleaseDate},
			Distributor: movies[i].Distributor,
			Budget:      movies[i].Budget,
			MpaRating:   movies[i].MpaRating,
			Id:          movies[i].ID,
		}
		response, err := h.cli.GetMovieBoxOfficeWithResponse(c, &openapi.GetMovieBoxOfficeParams{Title: "线些西周"})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if err = json.Unmarshal(response.Body, &data.Items[i].BoxOffice); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
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

	if err := c.BindJSON(&m); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateMovie(&m); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, m)
}

func (h *MovieHandler) GetMoviesTitleRating(c *gin.Context, title string) {
	fmt.Println("GetMoviesTitleRating")
}

func (h *MovieHandler) PostMoviesTitleRatings(c *gin.Context, title string) {
	fmt.Println("PostMoviesTitleRatings")
}
