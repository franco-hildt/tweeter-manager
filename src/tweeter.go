package main

import (
	"github.com/abiosoft/ishell"
	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/domain"
	"github.com/franco-hildt/tweeter-manager/tweeter-manager/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")
	tweetManager := service.NewTweetManager()

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your tweet: ")

			text := c.ReadLine()
			tweet := domain.NewTextTweet("fhildt", text)

			tweetManager.PublishTweet(tweet)

			c.Print("Tweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showLastTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetLastTweet()
			c.Println(tweet.PrintableTweet())

			//printTweet(tweet, c)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := tweetManager.GetTweets()

			//c.Println(tweet)
			for _, t := range tweet {
				printTweet(t, c)
			}

			return
		},
	})

	shell.Run()

}

func printTweet(t domain.Tweet, c *ishell.Context) {
	// c.Println(strconv.Itoa(t.Id) + ". " + t.Text)
	// c.Println("user:" + t.User + "  " + "date:" + t.Date.Format("2006-01-02 15:04:05")) //"2006-01-02 15:04:05"
	c.Println(t.PrintableTweet())
}
