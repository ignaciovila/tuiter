package tweetManager

type QuoteTweet struct {
	TextTweet
	Tweet
}

func (tweet QuoteTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text + "\n" +
		"Quote: " + "@" + tweet.Tweet.GetUser() + ": " + tweet.Tweet.GetText()
}