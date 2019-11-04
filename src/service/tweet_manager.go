package service

import (
	"fmt"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
)

type TweetManager struct {
	Tweets       []domain.Tweet
	TweetsByUser map[string][]domain.Tweet
	LastId       int
}

func NewTweetManager() TweetManager {
	var this TweetManager
	this.Tweets = make([]domain.Tweet, 0)
	this.LastId = 0
	this.TweetsByUser = make(map[string][]domain.Tweet)

	return this
}

func (this *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {

	if tweet.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	} else if len(tweet.GetText()) > 140 {
		return 0, fmt.Errorf("text exceeds 140 characters")
	} else if len(tweet.GetText()) == 0 {
		return 0, fmt.Errorf("text is required")
	} else {
		this.LastId++
		tweet.SetId(this.LastId)
		this.Tweets = append(this.Tweets, tweet)
		this.TweetsByUser[tweet.GetUser()] = append(this.TweetsByUser[tweet.GetUser()], tweet)
		return this.LastId, nil
	}
}

func (this *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return this.TweetsByUser[user]
}

func (this *TweetManager) GetTweets() []domain.Tweet {
	return this.Tweets
}

func (this *TweetManager) GetLastTweet() domain.Tweet {
	return this.Tweets[len(this.Tweets)-1]
}

func (this *TweetManager) GetTweetById(id int) domain.Tweet {
	for _, t := range this.Tweets {
		if t.GetId() == id {
			return t
		}
	}
	return nil
}

func (this *TweetManager) CountTweetsByUser(user string) (cont int) {

	return len(this.TweetsByUser[user])
}
