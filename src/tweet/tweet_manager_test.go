package tweet_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/tweet"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	var twt *tweet.Tweet
	user := "ignaciovila"
	text := "This is my third tweet"
	twt = tweet.NewTweet(user, text)

	// Operation
	tweet.PublishTweet(twt)

	// Validation
	publishedTweet := tweet.GetTweets()[0]
	if publishedTweet.User != user ||
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var twt *tweet.Tweet
	var user string
	text := "This is my first tweet"
	twt = tweet.NewTweet(user, text)

	var err error
	err = tweet.PublishTweet(twt)

	if err == nil || err.Error() != "user is required" {
		t.Error("Expected error is: user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var twt *tweet.Tweet
	user := "ignaciovila"
	var text string
	twt = tweet.NewTweet(user, text)

	var err error
	err = tweet.PublishTweet(twt)

	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is: text is required")
	}
}

func TestTweetThatExceedes140CharactersIsNotPublished(t *testing.T) {
	var twt *tweet.Tweet
	user := "ignaciovila"
	text := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	twt = tweet.NewTweet(user, text)

	var err error
	err = tweet.PublishTweet(twt)

	if err == nil || err.Error() != "max length is 140" {
		t.Error("Expected error is: max length is 140")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	tweet.InitializateService()
	var twt, secondTweet *tweet.Tweet
	user := "ignaciovila"
	text := "This is my first tweet"
	twt = tweet.NewTweet(user, text)
	secondTweet = tweet.NewTweet(user, text)

	tweet.PublishTweet(twt)
	tweet.PublishTweet(secondTweet)

	publishedTweets := tweet.GetTweets()
	if len(publishedTweets) != 2 {
		t.Error("Unexpected size of tweet list")
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, user, text)
	isValidTweet(t, secondPublishedTweet, user, text)
}

func isValidTweet(t *testing.T, twt *tweet.Tweet, user string, text string) {
	if twt.User != user {
		t.Error("invalid user")
	}

	if twt.Text != text {
		t.Error("invalid text")
	}
}
