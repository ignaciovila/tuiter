package service_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/domain"
	"github.com/ignaciovila/tuiter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	var tweet *domain.Tweet
	user := "ignaciovila"
	text := "This is my third tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user ||
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	var user string
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil || err.Error() != "user is required" {
		t.Error("Expected error is: user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "ignaciovila"
	var text string
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is: text is required")
	}
}

func TestTweetThatExceedes140CharactersIsNotPublished(t *testing.T) {
	var tweet *domain.Tweet
	user := "ignaciovila"
	text := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	tweet = domain.NewTweet(user, text)

	var err error
	err = service.PublishTweet(tweet)

	if err == nil || err.Error() != "max length is 140" {
		t.Error("Expected error is: max length is 140")
	}
}
