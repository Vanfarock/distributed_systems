package main

import (
	"fmt"
	"time"

	"github.com/Vanfarock/twitter_distributed/tweet_writer/tweeter"
	"github.com/joho/godotenv"
)

func main() {
	// count := 0
	// r := gin.Default()

	// r.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, count)
	// })

	// r.GET("/sum", func(ctx *gin.Context) {
	// 	count += 1
	// })

	godotenv.Load(".env")

	err := Migrate()
	if err != nil {
		fmt.Println(err)
	}

	tweet, _ := tweeter.NewTweet(1, "Teste", time.Now())
	repository, _ := tweeter.NewTweetRepository()
	service := tweeter.NewTweetService(repository)
	err = service.Tweet(tweet)
	if err != nil {
		fmt.Println(err)
	}

	// r.Run(":8080")
}
