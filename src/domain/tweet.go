package domain

import "time"

type Tweet struct {
	Text string
	User string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	p := time.Now()
	tweet := Tweet{text, user, &p}

	return &tweet
}
