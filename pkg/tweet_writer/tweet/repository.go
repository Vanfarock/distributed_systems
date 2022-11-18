package tweet

type TweetRepository interface {
	GetByUser(userId int64) ([]*Tweet, error)
	Save(tweet Tweet) error
}

type DefaultTweetRepository struct{}

func NewTweetRepository() (*DefaultTweetRepository, error) {
	return &DefaultTweetRepository{}, nil
}

func (DefaultTweetRepository) GetByUser(userId int64) ([]*Tweet, error) {
	return nil, nil
}

func (DefaultTweetRepository) Save(tweet Tweet) error {
	return nil
}
