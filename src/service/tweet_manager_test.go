package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	assert.Equal(t, user, publishedTweet.User, "Should be equal")
	assert.Equal(t, text, publishedTweet.Text, "Should be equal")
	assert.NotNil(t, publishedTweet.Date, "Should NOT be equal")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err, "Should be nil")
	assert.Equal(t, "user is required", err.Error(), "Should be nil")
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

}

func TestTweetWichEwxceeding140CharactersIsNotPublished(t *testing.T) {

}
