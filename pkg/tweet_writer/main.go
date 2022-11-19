package main

import (
	"fmt"

	"github.com/Vanfarock/twitter_distributed/tweet_writer/tweeter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	err := Migrate()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()
	r.POST("/", tweeter.CreateTweet)
	r.Run(":8080")
}
