package main

import (
	"UrlShortener/internal/adapter/input_port"
	"UrlShortener/internal/adapter/output_port"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := output_port.NewShortUrlInfoMockRepository(nil)
	controller := input_port.NewShortUrlInfoController(repository)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, Url Shortener!")
	})

	r.POST("/api/short_url", controller.CreateShortUrlCode)
	r.GET("/api/short_url", controller.GetOGUrl)
	r.Run(":8000")
}
