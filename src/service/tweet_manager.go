package service

import (
	"fmt"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
)

var Tweet []domain.Tweet

func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return fmt.Errorf("user is required")
	} else if len(tweet.Text) > 140 {
		return fmt.Errorf("text exceeds 140 characters")
	} else if len(tweet.Text) == 0 {
		return fmt.Errorf("text is required")
	} else {
		Tweet = append(Tweet, *tweet)
	}
	return nil
}

func GetTweets() []domain.Tweet {
	return Tweet
}

func GetLastTweet() domain.Tweet {
	return Tweet[len(Tweet)-1]
}

func InitializeService() {
	Tweet = make([]domain.Tweet, 0)
}
