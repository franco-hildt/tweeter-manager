package domain

import (
	"fmt"
	"time"
)

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetId() int
	SetId(int)
	PrintableTweet() string
}

type TextTweet struct {
	Text string
	User string
	Date *time.Time
	Id   int
}

type ImageTweet struct {
	TextTweet
	Link string
}

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func NewTextTweet(user string, text string) *TextTweet {
	p := time.Now()
	tweet := TextTweet{text, user, &p, 0}

	return &tweet
}

func (tweet *TextTweet) GetUser() string {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) GetId() int {
	return tweet.Id
}

func (tweet *TextTweet) SetId(id int) {
	tweet.Id = id
}

func (this *TextTweet) PrintableTweet() string {
	return "@" + this.User + ": " + this.Text
}

func (this *TextTweet) String() string {
	return fmt.Sprintf("@%v: %v", this.User, this.Text)
}

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {
	tweetAux := NewTextTweet(user, text)

	tweet := QuoteTweet{TextTweet: *tweetAux, QuotedTweet: quotedTweet}
	return &tweet
}

func (tweet *QuoteTweet) GetUser() string {
	return tweet.User
}

func (tweet *QuoteTweet) GetText() string {
	return tweet.Text
}

func (tweet *QuoteTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *QuoteTweet) GetId() int {
	return tweet.Id
}

func (tweet *QuoteTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf(`@%s: %s "%s"`, tweet.User, tweet.Text, tweet.QuotedTweet)
}

func (tweet *QuoteTweet) String() string {
	return tweet.PrintableTweet()
}

func NewImageTweet(user string, text string, link string) *ImageTweet {
	tweetAux := NewTextTweet(user, text)
	tweet := ImageTweet{TextTweet: *tweetAux, Link: link}

	return &tweet
}

func (tweet *ImageTweet) GetUser() string {
	return tweet.User
}

func (tweet *ImageTweet) GetText() string {
	return tweet.Text
}

func (tweet *ImageTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *ImageTweet) GetId() int {
	return tweet.Id
}

func (tweet *ImageTweet) SetId(id int) {
	tweet.Id = id
}

func (tweet *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", tweet.User, tweet.Text, tweet.Link)
}

func (tweet *ImageTweet) String() string {
	return tweet.PrintableTweet()
}
