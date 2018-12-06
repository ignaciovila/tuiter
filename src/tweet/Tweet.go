package tweetManager

import "time"

type Tweet interface {
	String() string
	GetUser() string
	GetText() string
	GetDate() *time.Time
}