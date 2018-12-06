package tweetManager

type ImageTweet struct {
	TextTweet
	Url string
}

func (tweet ImageTweet) String() string {
	return "@" + tweet.User + ": " + tweet.Text + "\n" +
		"URL: " + tweet.Url
}
