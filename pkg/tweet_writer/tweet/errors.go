package tweet

import "fmt"

type TweetCreationError struct {
	Err error
}

func (e *TweetCreationError) Error() string {
	return fmt.Sprintf("Error while creating a tweet: %d", e.Err)
}
