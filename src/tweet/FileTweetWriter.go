package tweetManager

import (
	"fmt"
	"os"
)

type FileTweetWriter struct {
	file *os.File
}

func (writer *FileTweetWriter) WriteTweet(tweet Tweet, c chan int) {
	fmt.Fprintf(writer.file, tweet.String())
}

func (writer *FileTweetWriter) getTweets() []Tweet{
	//TODO implement
	return nil
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.Create("tweets.txt")
	return &FileTweetWriter{file}
}
