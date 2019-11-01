package service

import (
	"fmt"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
)

var Tweets []domain.Tweet
var TweetsByUser map[string][]domain.Tweet
var lastId int = 0

func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, fmt.Errorf("user is required")
	} else if len(tweet.Text) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	} else if len(tweet.Text) == 0 {
		return 0, fmt.Errorf("text is required")
	} else {
		lastId++
		tweet.Id = lastId
		Tweets = append(Tweets, *tweet)
		TweetsByUser[tweet.User] = append(TweetsByUser[tweet.User], *tweet)
		return lastId, nil
	}
}

func GetTweetsByUser(user string) []domain.Tweet {
	return TweetsByUser[user]
}

func GetTweets() []domain.Tweet {
	return Tweets
}

func GetLastTweet() domain.Tweet {
	return Tweets[len(Tweets)-1]
}

func InitializeService() {
	Tweets = make([]domain.Tweet, 0)
	lastId = 0
	TweetsByUser = make(map[string][]domain.Tweet)
}

func GetTweetById(id int) *domain.Tweet {
	for _, t := range Tweets {
		if t.Id == id {
			return &t
		}
	}
	return nil
}

func CountTweetsByUser(user string) (cont int) {
	cont = 0
	for _, t := range Tweets {
		if t.User == user {
			cont++
		}
	}
	return
}
