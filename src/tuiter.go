package main

import (
	"github.com/abiosoft/ishell"
	"github.com/ignaciovila/tuiter/src/api"
	"github.com/ignaciovila/tuiter/src/tweet"
	"github.com/ignaciovila/tuiter/src/user"
)

func main() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	service := tweet.NewTweetManager(tweet.NewMemoryTweetWriter())

	server := api.NewGinServer(service)

	server.StartGinServer()

	user.AddUser(user.NewUser("name", "mail", "nacho", "pass"))
	service.PublishTweet(tweet.NewTextTweet("nacho", "hola mundo"))

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			tweet := tweet.NewTextTweet("some user", text)

			service.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweets()

			c.Println(tweet)

			return
		},
	})

	shell.Run()

}
