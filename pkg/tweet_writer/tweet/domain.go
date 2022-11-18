package tweet

import (
	"errors"
	"time"
)

type Tweet struct {
	Id             int64
	AuthorId       int64
	Content        string
	Likes          int32
	Retweets       int32
	DateTime       time.Time
	ReferenceTweet *int64
}

func NewTweet(authorId int64, content string, dateTime time.Time) (*Tweet, error) {
	if len(content) > getTweetMaxChars() {
		return nil, &TweetCreationError{
			Err: errors.New("Tweet's contents are greater than maximum allowed"),
		}
	}

	tweet := Tweet{
		AuthorId:       authorId,
		Content:        content,
		Likes:          0,
		Retweets:       0,
		DateTime:       dateTime,
		ReferenceTweet: nil,
	}
	return &tweet, nil
}

func getTweetMaxChars() int {
	return 256
}
