package tweet

import "time"

type ImageTweet struct {
	TextTweet
	Url string
}

func (tweet *ImageTweet) String() string {
	return "@" + tweet.GetUser() + ": " + tweet.GetText() + " - " +
		"URL: " + tweet.Url
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	now := time.Now()
	return &ImageTweet{TextTweet{user, text, &now}, url}
}