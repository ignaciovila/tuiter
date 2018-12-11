package api_test

import (
	"github.com/ignaciovila/tuiter/src/api"
	"github.com/ignaciovila/tuiter/src/tweet"
	"testing"
)

func TestGinServer(t *testing.T) {
	server := api.NewGinServer(tweet.NewTweetManager(tweet.NewMemoryTweetWriter()))

	server.StartGinServer()
}


