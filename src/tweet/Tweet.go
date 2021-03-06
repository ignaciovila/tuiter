package tweetManager

import (
	"time"
)

type Tweet struct {
	User string
	Text string
	Date *time.Time
}

func (tweet Tweet) PrintableTweet() string {
	return "@" + tweet.User + ": " + tweet.Text
}

func NewTweet(user string, text string) *Tweet {
	now := time.Now()
	return &Tweet{user, text, &now}
}
