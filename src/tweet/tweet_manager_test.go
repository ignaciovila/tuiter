package tweetManager_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/tweet"
	"github.com/ignaciovila/tuiter/src/user"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	tweetManager.NewTweetManager()
	var twt *tweetManager.Tweet
	usr := "ignaciovila"
	text := "This is my third tweet"
	twt = tweetManager.NewTweet(usr, text)

	usr1 := userManager.NewUser("name", "mail", "ignaciovila", "password")
	userManager.AddUser(usr1)

	// Operation
	_, err := tweetManager.PublishTweet(twt)

	// Validation
	if err!= nil && err.Error() == "user is required" {
		t.Error("user is required")
	}

	publishedTweet := tweetManager.GetTweets()[0]
	if publishedTweet.User != usr ||
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", usr, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date cannot be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var twt *tweetManager.Tweet
	var user string
	text := "This is my first tweet"
	twt = tweetManager.NewTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(twt)

	if err == nil || err.Error() != "user is required" {
		t.Error("Expected error is: user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var twt *tweetManager.Tweet
	user := "ignaciovila"
	var text string
	twt = tweetManager.NewTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(twt)

	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is: text is required")
	}
}

func TestTweetThatExceedes140CharactersIsNotPublished(t *testing.T) {
	var twt *tweetManager.Tweet
	user := "ignaciovila"
	text := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	twt = tweetManager.NewTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(twt)

	if err == nil || err.Error() != "max length is 140" {
		t.Error("Expected error is: max length is 140")
	}
}

func TestTweetWithInvalidUserIsNotPublished(t *testing.T) {
	var twt *tweetManager.Tweet
	user := "manolo"
	text := "This is my first tweet"
	twt = tweetManager.NewTweet(user, text)

	var err error
	_, err = tweetManager.PublishTweet(twt)

	if err == nil || err.Error() != "invalid user" {
		t.Error("Expected error is: invalid user")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	tweetManager.NewTweetManager()
	var twt, secondTweet *tweetManager.Tweet
	user := "ignaciovila"
	text1 := "This is my first tweet"
	text2 := "This is my second tweet"
	twt = tweetManager.NewTweet(user, text1)
	secondTweet = tweetManager.NewTweet(user, text2)

	tweetManager.PublishTweet(twt)
	tweetManager.PublishTweet(secondTweet)

	publishedTweets := tweetManager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Error("Unexpected size of tweet list")
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, user, text1)
	isValidTweet(t, secondPublishedTweet, user, text2)
}

func TestCanRetrieveTweetById(t *testing.T) {
	tweetManager.NewTweetManager()
	var twt *tweetManager.Tweet
	var id int
	user := "ignaciovila"
	text := "This is my third tweet"
	twt = tweetManager.NewTweet(user, text)

	id, _ = tweetManager.PublishTweet(twt)

	publishedTweet := tweetManager.GetTweetById(id)
	isValidTweet(t, publishedTweet, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager.NewTweetManager()
	var twt, secondTweet, thirdTweet *tweetManager.Tweet
	usr := "grupoesfera"
	anotherUser := "nick"

	usr1 := userManager.NewUser("name", "mail", "grupoesfera", "password")
	usr2 := userManager.NewUser("name", "mail", "nick", "password")

	userManager.AddUser(usr1)
	userManager.AddUser(usr2)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	twt = tweetManager.NewTweet(usr, text)
	secondTweet = tweetManager.NewTweet(usr, secondText)
	thirdTweet = tweetManager.NewTweet(anotherUser, text)
	tweetManager.PublishTweet(twt)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(usr)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager.NewTweetManager()
	var twt, secondTweet, thirdTweet *tweetManager.Tweet
	usr := "grupocubo"
	anotherUser := "nick"

	usr1 := userManager.NewUser("name", "mail", "grupocubo", "password")
	usr2 := userManager.NewUser("name", "mail", "nick", "password")

	userManager.AddUser(usr1)
	userManager.AddUser(usr2)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	twt = tweetManager.NewTweet(usr, text)
	secondTweet = tweetManager.NewTweet(usr, secondText)
	thirdTweet = tweetManager.NewTweet(anotherUser, text)
	tweetManager.PublishTweet(twt)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetByUser(usr)

	// Validation
	if len(tweets) !=2 {
		t.Error("invalid amount of tweets for user")
	}
}

func isValidTweet(t *testing.T, twt *tweetManager.Tweet, user string, text string) {
	if twt.User != user {
		t.Error("invalid user")
	}

	if twt.Text != text {
		t.Error("invalid text")
	}
}
