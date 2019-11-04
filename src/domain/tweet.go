package domain

import (
	"fmt"
	"time"
)

type Tweet struct {
	Text string
	User string
	Date *time.Time
	Id   int
}

func NewTweet(user string, text string) *Tweet {
	p := time.Now()
	tweet := Tweet{text, user, &p, 0}

	return &tweet
}

func (this *Tweet) PrintableTweet() string {
	return "@" + this.User + ": " + this.Text
}

func (this *Tweet) String() string {
	return fmt.Sprintf("@%v: %v", this.User, this.Text)
}
