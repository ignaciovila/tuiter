package tweetManager_test

import (
	"testing"

	"github.com/ignaciovila/tuiter/src/tweet"
	"github.com/ignaciovila/tuiter/src/user"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	// Initialization
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)
	var twt *tweetManager.TextTweet
	usr := "ignaciovila"
	text := "This is my third tweet"
	twt = tweetManager.NewTextTweet(usr, text)

	usr1 := userManager.NewUser("name", "mail", "ignaciovila", "password")
	userManager.AddUser(usr1)

	// Operation
	_, err := manager.PublishTweet(twt)

	// Validation
	if err!= nil && err.Error() == "user is required" {
		t.Error("user is required")
	}

	publishedTweet := manager.GetTweets()[0]

	if publishedTweet.GetUser() != usr ||
		publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", usr, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Error("Expected date cannot be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	var twt *tweetManager.TextTweet
	var user string
	text := "This is my first tweet"
	twt = tweetManager.NewTextTweet(user, text)
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var err error
	_, err = manager.PublishTweet(twt)

	if err == nil || err.Error() != "user is required" {
		t.Error("Expected error is: user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	var twt *tweetManager.TextTweet
	user := "ignaciovila"
	var text string
	twt = tweetManager.NewTextTweet(user, text)
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var err error
	_, err = manager.PublishTweet(twt)

	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is: text is required")
	}
}

func TestTweetThatExceedes140CharactersIsNotPublished(t *testing.T) {
	var twt *tweetManager.TextTweet
	user := "ignaciovila"
	text := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	twt = tweetManager.NewTextTweet(user, text)
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var err error
	_, err = manager.PublishTweet(twt)

	if err == nil || err.Error() != "max length is 140" {
		t.Error("Expected error is: max length is 140")
	}
}

func TestTweetWithInvalidUserIsNotPublished(t *testing.T) {
	var twt *tweetManager.TextTweet
	user := "manolo"
	text := "This is my first tweet"
	twt = tweetManager.NewTextTweet(user, text)
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var err error
	_, err = manager.PublishTweet(twt)

	if err == nil || err.Error() != "invalid user" {
		t.Error("Expected error is: invalid user")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)
	var twt, secondTweet *tweetManager.TextTweet
	user := "ignaciovila"
	text1 := "This is my first tweet"
	text2 := "This is my second tweet"
	twt = tweetManager.NewTextTweet(user, text1)
	secondTweet = tweetManager.NewTextTweet(user, text2)

	manager.PublishTweet(twt)
	manager.PublishTweet(secondTweet)

	publishedTweets := manager.GetTweets()
	if len(publishedTweets) != 2 {
		t.Error("Unexpected size of tweet list")
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	isValidTweet(t, firstPublishedTweet, user, text1)
	isValidTweet(t, secondPublishedTweet, user, text2)
}

func TestCanRetrieveTweetById(t *testing.T) {
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)
	var twt *tweetManager.TextTweet
	var id int
	user := "ignaciovila"
	text := "This is my third tweet"
	twt = tweetManager.NewTextTweet(user, text)

	id, _ = manager.PublishTweet(twt)

	publishedTweet := manager.GetTweetById(id)
	isValidTweet(t, publishedTweet, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)
	var twt, secondTweet, thirdTweet *tweetManager.TextTweet
	usr := "grupoesfera"
	anotherUser := "nick"

	usr1 := userManager.NewUser("name", "mail", "grupoesfera", "password")
	usr2 := userManager.NewUser("name", "mail", "nick", "password")

	userManager.AddUser(usr1)
	userManager.AddUser(usr2)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	twt = tweetManager.NewTextTweet(usr, text)
	secondTweet = tweetManager.NewTextTweet(usr, secondText)
	thirdTweet = tweetManager.NewTextTweet(anotherUser, text)
	manager.PublishTweet(twt)
	manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	// Operation
	count := manager.CountTweetsByUser(usr)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)
	var twt, secondTweet, thirdTweet *tweetManager.TextTweet
	usr := "grupocubo"
	anotherUser := "nick"

	usr1 := userManager.NewUser("name", "mail", "grupocubo", "password")
	usr2 := userManager.NewUser("name", "mail", "nick", "password")

	userManager.AddUser(usr1)
	userManager.AddUser(usr2)

	text := "This is my first tweet"
	secondText := "This is my second tweet"
	twt = tweetManager.NewTextTweet(usr, text)
	secondTweet = tweetManager.NewTextTweet(usr, secondText)
	thirdTweet = tweetManager.NewTextTweet(anotherUser, text)
	manager.PublishTweet(twt)
	manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	// Operation
	tweets := manager.GetTweetByUser(usr)

	// Validation
	if len(tweets) !=2 {
		t.Error("invalid amount of tweets for user")
	}
}

func TestPrintableTweet(t *testing.T) {
	twt := tweetManager.NewTextTweet("nacho", "este tuit se va a ver bonito")

	niceTweet := twt.String()

	if niceTweet != "@nacho: este tuit se va a ver bonito" {
		t.Error("el tuit no se vio bonito")
	}
}

func TestImageTweet(t *testing.T) {
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var twt *tweetManager.ImageTweet
	var id int
	userText := "ignaciovila"
	text := "soy un tuit con imagen!"

	user := userManager.NewUser("name", "mail", "ignaciovila", "password")
	userManager.AddUser(user)

	twt = tweetManager.NewImageTweet(userText, text, "http://myserver.com/randomImage")

	id, _ = manager.PublishTweet(twt)

	publishedTweet := manager.GetTweetById(id)
	isValidTweet(t, publishedTweet, userText, text)
	imageTweet, _ := publishedTweet.(*tweetManager.ImageTweet)
	if imageTweet.Url != twt.Url {
		t.Error("invalid url")
	}
}

func TestQuoteTweet(t *testing.T) {
	writer := tweetManager.NewMemoryTweetWriter()
	manager := tweetManager.NewTweetManager(writer)

	var qTwt *tweetManager.QuoteTweet
	var twt *tweetManager.TextTweet
	var id int
	userText := "ignaciovila"
	text := "soy un tuit embarazado!"

	user := userManager.NewUser("name", "mail", "ignaciovila", "password")
	userManager.AddUser(user)


	twt = tweetManager.NewTextTweet("juan carlos", "me estan citando!")
	qTwt = tweetManager.NewQuoteTweet(userText, text, twt)

	id, _ = manager.PublishTweet(qTwt)

	publishedTweet := manager.GetTweetById(id)
	isValidTweet(t, publishedTweet, userText, text)
	quoteTweet, _ := publishedTweet.(*tweetManager.QuoteTweet)
	if quoteTweet.QuotedTweet.GetText() != twt.Text {
		t.Error("invalid quote")
	}
}

func TestTweetWriter(t *testing.T) {

}

func isValidTweet(t *testing.T, twt tweetManager.Tweet, user string, text string) {
	if twt.GetUser() != user {
		t.Error("invalid user: " + twt.GetUser())
	}

	if twt.GetText() != text {
		t.Error("invalid text: " + twt.GetText())
	}
}
