package tweetManager

type TweetWriter interface {
	WriteTweet(tweet Tweet)
	getTweets()	[]Tweet
}