package tweet

type TweetWriter interface {
	WriteTweet(tweet Tweet)
	getTweets()	[]Tweet
}