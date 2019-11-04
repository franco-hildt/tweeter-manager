package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
)

func TestCanGetAPrintableTweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"
	assert.Equal(t, expectedText, text, "The texts should be equal")
}

func TestCanGetAStringFromATweet(t *testing.T) {

	// Initialization
	tweet := domain.NewTweet("grupoesfera", "This is my tweet")

	// Operation
	text := tweet.String()

	// Validation
	expectedText := "@grupoesfera: This is my tweet"

	assert.Equal(t, text, expectedText)

	// if text != expectedText {
	//     t.Errorf("The expected text is %s but was %s", expectedText, text)
	// }
}
