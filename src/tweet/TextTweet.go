package tweet

import (
	"time"
)

type TextTweet struct {
	user string
	text string
	date *time.Time
}

func NewTextTweet(user string, text string) *TextTweet {
	now := time.Now()
	return &TextTweet{user, text, &now}
}

func (tweet *TextTweet) String() string {
	return "@" + tweet.user + ": " + tweet.text
}

func (tweet *TextTweet) GetUser() string {
	return tweet.user
}

func (tweet *TextTweet) GetText() string {
	return tweet.text
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.date
}