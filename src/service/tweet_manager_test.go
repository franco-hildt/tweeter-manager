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
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet := domain.NewTextTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetLastTweet()

	assert.Equal(t, user, publishedTweet.GetUser(), "Should be equal")
	assert.Equal(t, text, publishedTweet.GetText(), "Should be equal")
	assert.NotNil(t, publishedTweet.GetDate(), "Should NOT be equal")
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	var user string
	text := "This is my first tweet"

	tweet := domain.NewTextTweet(user, text)

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
	var tweet, secondTweet domain.Tweet // Fill the tweets with data
	text1 := "Primer tweet"
	text2 := "Segundo tweet"
	user := "fhildt"
	tweet = domain.NewTextTweet(user, text1)
	secondTweet = domain.NewTextTweet(user, text2)

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()
	assert.Equal(t, 2, len(publishedTweets), "Should be the same len")

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	assert.Equal(t, firstPublishedTweet.GetText(), text1, "Should be equal")
	assert.Equal(t, firstPublishedTweet.GetUser(), user, "Should be equal")

	assert.Equal(t, secondPublishedTweet.GetText(), text2, "Should be equal")
	assert.Equal(t, secondPublishedTweet.GetUser(), user, "Should be equal")

}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet domain.Tweet

	user := "grupoesfera"
	var text string

	tweet = domain.NewTextTweet(user, text)

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
	var tweet domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTextTweet(user, text)

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

	var tweet domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	assert.Equal(t, publishedTweet.GetText(), text)
	assert.Equal(t, publishedTweet.GetUser(), user, "Should be equal")
	assert.Equal(t, publishedTweet.GetId(), id, "Should be equal")
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	var tweet, secondTweet, thirdTweet domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)
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

	var tweet, secondTweet, thirdTweet domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

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

	assert.Equal(t, firstPublishedTweet.GetUser(), user, "Should be equal")
	assert.Equal(t, firstPublishedTweet.GetText(), text, "Should be equal")
	assert.Equal(t, secondPublishedTweet.GetUser(), user, "Should be equal")
	assert.Equal(t, secondPublishedTweet.GetText(), secondText, "Should be equal")
}
