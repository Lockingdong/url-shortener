package main

import (
	"UrlShortener/internal/adapter/input_port"
	"UrlShortener/internal/adapter/output_port"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
)

func main() {

	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// if _, err := client.Ping(context.Background()).Result(); err != nil {
	// 	log.Fatal(err)
	// }
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
