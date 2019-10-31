package service_test

import (
	"testing"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/service"
	"github.com/stretchr/testify/assert"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet string = "This is my first tweet"

	service.PublishTweet(tweet)

	assert.Equal(t, tweet, service.GetTweet(), "They should be equal")
}
