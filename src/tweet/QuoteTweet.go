package tweet

import "time"

type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet
}

func (tweet *QuoteTweet) String() string {
	return "@" + tweet.GetUser() + ": " + tweet.GetText() + " - " +
		"Quote: " + "@" + tweet.QuotedTweet.GetUser() + ": " + tweet.QuotedTweet.GetText()
}

func NewQuoteTweet(user string, text string, tweet Tweet) *QuoteTweet {
	now := time.Now()
	return &QuoteTweet{TextTweet{user, text, &now}, tweet}
}