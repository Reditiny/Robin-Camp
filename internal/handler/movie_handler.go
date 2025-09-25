package handler

import (
	"assignment/config"
	"assignment/internal/client"
	"assignment/openapi"
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println("GetMovies")
	response, err := h.cli.GetMovieBoxOfficeWithResponse(c, &openapi.GetMovieBoxOfficeParams{Title: "线些西周"})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(response)
}

func (h *MovieHandler) PostMovies(c *gin.Context) {
	fmt.Println("PostMovies")
}

func (h *MovieHandler) GetMoviesTitleRating(c *gin.Context, title string) {
	fmt.Println("GetMoviesTitleRating")
}

func (h *MovieHandler) PostMoviesTitleRatings(c *gin.Context, title string) {
	fmt.Println("PostMoviesTitleRatings")
}
