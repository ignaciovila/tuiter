package tweet

import (
	"fmt"
	)

var tweets []*Tweet

func PublishTweet(t *Tweet) error {
	if len(t.User) < 1 {
		return fmt.Errorf("user is required")
	}

	if len(t.Text) < 1 {
		return fmt.Errorf("text is required")
	}

	if len(t.Text) > 140 {
		return fmt.Errorf("max length is 140")
	}

	tweets = append(tweets, t)

	return nil
}

func GetTweets() []*Tweet {
	return tweets
}

func InitializateService() {
	tweets = make([]*Tweet, 0)
}
