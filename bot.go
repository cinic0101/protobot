package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main()  {
	token := os.Getenv("slack-bot-token")

	if token == "" {
		fmt.Println("[ERROR] No Environment variable \"token\". ")
		return
	}

	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			if ev.BotID == "" {
				params := slack.PostMessageParameters{
					AsUser: true,
				}
				attachment := slack.Attachment{
					Text: ev.Text,
				}
				params.Attachments = []slack.Attachment{attachment}
				api.PostMessage(ev.Channel, "Response", params)
			}

		}
	}
	
}
