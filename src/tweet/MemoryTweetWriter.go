package tweet

type MemoryTweetWriter struct {
	tweets []Tweet
}

func (writer *MemoryTweetWriter) WriteTweet(tweet Tweet) {
	writer.tweets = append(writer.tweets, tweet)
}

func (writer *MemoryTweetWriter) getTweets() []Tweet {
	return writer.tweets
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{}
}