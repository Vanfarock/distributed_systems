package tweeter

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Tweet struct {
	gorm.Model
	UserId  int64  `gorm:"index"`
	Content string `gorm:"type:varchar(255)" json:"content"`
}

func NewTweet(authorId int64, content string, dateTime time.Time) (*Tweet, error) {
	if len(content) > getTweetMaxChars() {
		return nil, NewTweetCreationError(errors.New("Tweet's contents are greater than maximum allowed"))
	}

	tweet := Tweet{
		UserId:  authorId,
		Content: content,
	}
	return &tweet, nil
}

func getTweetMaxChars() int {
	return 256
}
