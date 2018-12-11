package tweet

import (
	"os"
)

type FileTweetWriter struct {
	file *os.File
}

func (writer *FileTweetWriter) WriteTweet(tweet Tweet) {
	writer.file.WriteString(tweet.String())
}

func (writer *FileTweetWriter) getTweets() []Tweet{
	//TODO implement
	return nil
}

func NewFileTweetWriter() *FileTweetWriter {
	file, _ := os.Create("tweets.txt")
	return &FileTweetWriter{file}
}
