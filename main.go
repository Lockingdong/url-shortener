package main

import (
	"UrlShortener/internal/adapter/input_port"
	"UrlShortener/internal/adapter/output_port"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// repo := output_port.NewUrlInfoRepository(client)

	repo := output_port.NewUrlInfoMockRepository(nil)

	controller := input_port.NewUrlShortenerController(repo)

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello, Url Shortener!")
	})

	r.POST("/api/short_url", controller.CreateUrlCode)
	r.GET("/api/short_url", controller.GetUrl)
	r.Run(":8000")
}
