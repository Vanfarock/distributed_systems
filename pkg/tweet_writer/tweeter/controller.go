package tweeter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTweet(c *gin.Context) {
	tweetRepository, err := NewTweetRepository()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	service := NewTweetService(tweetRepository)

	var tweet *Tweet
	err = c.ShouldBind(&tweet)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = service.Tweet(tweet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, tweet)
}
