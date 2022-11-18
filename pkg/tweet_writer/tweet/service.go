package tweet

import "errors"

type TweetService interface {
	Tweet(tweet Tweet) error
}

type DefaultTweetService struct {
	tweetRepository TweetRepository
}

func NewTweetService(tweetRepository TweetRepository) DefaultTweetService {
	return DefaultTweetService{
		tweetRepository: tweetRepository,
	}
}

func (s DefaultTweetService) Tweet(tweet Tweet) error {
	if len(tweet.Content) > getTweetMaxChars() {
		return &TweetCreationError{
			Err: errors.New("Tweet's contents are greater than maximum allowed"),
		}
	}

	return s.tweetRepository.Save(tweet)
}
