package main

import (
	"github.com/Vanfarock/twitter_distributed/tweet_writer/db"
	"github.com/Vanfarock/twitter_distributed/tweet_writer/tweeter"
)

func Migrate() error {
	db, err := db.Open()
	if err != nil {
		return err
	}

	db.AutoMigrate(&tweeter.Tweet{})

	return nil
}
