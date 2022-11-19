package tweeter

import (
	"errors"
	"fmt"
)

type TweetCreationError struct {
	Err error
}

func (e *TweetCreationError) Error() string {
	return fmt.Sprintf("Error while creating a tweet: %d", e.Err)
}

func NewTweetCreationError(err error) *TweetCreationError {
	return &TweetCreationError{
		Err: errors.New("Tweet's contents are greater than maximum allowed"),
	}
}
