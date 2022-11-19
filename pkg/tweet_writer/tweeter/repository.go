package tweeter

import "github.com/Vanfarock/twitter_distributed/tweet_writer/db"

type TweetRepository interface {
	GetByUser(userId int64, limit int) ([]*Tweet, error)
	Create(tweet *Tweet) error
}

type DefaultTweetRepository struct{}

func NewTweetRepository() (*DefaultTweetRepository, error) {
	return &DefaultTweetRepository{}, nil
}

func (DefaultTweetRepository) GetByUser(userId int64, limit int) ([]*Tweet, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}

	var tweets []*Tweet
	db.Where("userId = ?", userId).
		Order("id desc").
		Limit(limit).
		Find(&tweets)

	return tweets, nil
}

func (DefaultTweetRepository) Create(tweet *Tweet) error {
	db, err := db.Open()
	if err != nil {
		return err
	}

	return db.Create(&tweet).Error
}
