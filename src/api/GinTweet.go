package api

type GinTweet struct {
	User string `json:"user"`
	Text string `json:"text"`
	Url string `json:"url"`
	Id string `json:"id"`
}