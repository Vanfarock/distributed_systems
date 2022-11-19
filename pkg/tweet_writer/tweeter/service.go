package tweeter

import "errors"

type TweetService interface {
	Tweet(tweet *Tweet) error
}

type DefaultTweetService struct {
	tweetRepository TweetRepository
}

func NewTweetService(tweetRepository TweetRepository) DefaultTweetService {
	return DefaultTweetService{
		tweetRepository: tweetRepository,
	}
}

func (s DefaultTweetService) Tweet(tweet *Tweet) error {
	if len(tweet.Content) > getTweetMaxChars() {
		return NewTweetCreationError(errors.New("Tweet's contents are greater than maximum allowed"))
	}

	return s.tweetRepository.Create(tweet)
}
