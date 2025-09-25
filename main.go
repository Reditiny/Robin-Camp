package main

import (
	"assignment/config"
	"assignment/internal/handler"
	"assignment/internal/repository"
	"assignment/openapi"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	if err := repository.InitDB(config.Conf.DbUrl); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	openapi.RegisterHandlers(r, handler.NewMovieHandler())

	if err := r.Run(fmt.Sprintf(":%s", config.Conf.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
