package tweetManager

import (
	"time"
)

type TextTweet struct {
	User string
	Text string
	Date *time.Time
}

func NewTextTweet(user string, text string) *TextTweet {
	now := time.Now()
	return &TextTweet{user, text, &now}
}

func (tweet *TextTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text
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