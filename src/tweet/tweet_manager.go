package tweet

import (
	"fmt"
	"github.com/ignaciovila/tuiter/src/user"
	)

type TweetManager struct {
	tweetsByUser map[string] []Tweet
	tweetWriter TweetWriter
}

func (manager *TweetManager) PublishTweet(t Tweet) (int, error) {

	if len(t.GetUser()) < 1 {
		return -1, fmt.Errorf("user is required")
	}

	if len(t.GetText()) < 1 {
		return -1, fmt.Errorf("text is required")
	}

	if len(t.GetText()) > 140 {
		return -1, fmt.Errorf("max length is 140")
	}

	if !user.ExistsUser(t.GetUser()) {
		return -1, fmt.Errorf("invalid user")
	}

	manager.tweetWriter.WriteTweet(t)

	currentList := manager.tweetsByUser[t.GetUser()]
	manager.tweetsByUser[t.GetUser()] = append(currentList, t)


	return len(manager.tweetWriter.getTweets()) -1, nil
}

func (manager *TweetManager) GetTweetById(index int) Tweet {
	return manager.tweetWriter.getTweets()[index]
}

func (manager *TweetManager) CountTweetsByUser(usr string) int {
	count := 0
	tweets := manager.tweetWriter.getTweets()

	for i := 0; i < len(tweets); i++ {
		if tweets[i].GetUser() == usr {
			count++
		}
	}

	return count
}

func (manager *TweetManager) GetTweetByUser(usr string) []Tweet {
	return manager.tweetsByUser[usr]
}

func (manager *TweetManager) GetTweets() []Tweet {
	return manager.tweetWriter.getTweets()
}

func NewTweetManager(writer TweetWriter)  *TweetManager {
	return &TweetManager{make(map[string] []Tweet), writer }
}
