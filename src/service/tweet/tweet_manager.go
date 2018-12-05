package service

import (
	"fmt"

	"github.com/ignaciovila/tuiter/src/domain"
)

var tweet domain.Tweet

// PublishTweet assign tweet to Tweet
func PublishTweet(t *domain.Tweet) error {
	if len(t.User) < 1 {
		return fmt.Errorf("user is required")
	}

	if len(t.Text) < 1 {
		return fmt.Errorf("text is required")
	}

	if len(t.Text) > 140 {
		return fmt.Errorf("max length is 140")
	}

	tweet = *t

	return nil
}

// GetTweet return current value
func GetTweet() domain.Tweet {
	return tweet
}
