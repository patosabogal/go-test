package main

import (
	"github.com/abiosoft/ishell"
	"github.com/patosabogal/go-test/src/domain"
	"github.com/patosabogal/go-test/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Username: ")

			var user string = c.ReadLine()

			c.Print("Write your tweet: ")

			var text string = c.ReadLine()

			var tweet = domain.NewTweet(user, text)

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
