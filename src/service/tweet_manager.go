package service

import (
	"fmt"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
)

var Tweet *domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	}
	Tweet = tweet
	return nil
}

func GetTweet() *domain.Tweet {
	return Tweet
}
