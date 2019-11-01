package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetLastTweet()

	assert.Equal(t, user, publishedTweet.User, "Should be equal")
	assert.Equal(t, text, publishedTweet.Text, "Should be equal")
	assert.NotNil(t, publishedTweet.Date, "Should NOT be equal")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err, "Should be nil")
	assert.Equal(t, "user is required", err.Error(), "Should be user is required")
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) { //
	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data
	text1 := "Primer tweet"
	text2 := "Segundo tweet"
	user := "fhildt"
	tweet = domain.NewTweet(user, text1)
	secondTweet = domain.NewTweet(user, text2)

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	assert.Equal(t, 2, len(publishedTweets), "Should be the same len")

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	assert.Equal(t, firstPublishedTweet.Text, text1, "Should be equal")
	assert.Equal(t, firstPublishedTweet.User, user, "Should be equal")

	assert.Equal(t, secondPublishedTweet.Text, text2, "Should be equal")
	assert.Equal(t, secondPublishedTweet.User, user, "Should be equal")

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err, "Should NOT be nil")

	assert.Equal(t, "text is required", err.Error(), "Should be the same")

}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	assert.NotNil(t, err, "Should NOT be nil")
	assert.Equal(t, "text exceeds 140 characters", err.Error(), "Should be equal")
}

func TestCanRetrieveTweetById(t *testing.T) { //

	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	assert.Equal(t, publishedTweet.Text, text, "Should be equal")
	assert.Equal(t, publishedTweet.User, user, "Should be equal")
	assert.Equal(t, publishedTweet.Id, id, "Should be equal")
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	// Operation
	count := tweetManager.CountTweetsByUser(user)
	// Validation

	assert.Equal(t, 2, count, "Should be equal")
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	// publish the 3 tweets
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	assert.Equal(t, 2, len(tweets), "Should be equal")
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet

	assert.Equal(t, firstPublishedTweet.User, user, "Should be equal")
	assert.Equal(t, firstPublishedTweet.Text, text, "Should be equal")
	assert.Equal(t, secondPublishedTweet.User, user, "Should be equal")
	assert.Equal(t, secondPublishedTweet.Text, secondText, "Should be equal")
}
