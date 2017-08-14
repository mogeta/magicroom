package main

import "github.com/nlopes/slack"

//If receive TV command. output to Irkit.
type TV struct {
	//relay.Irkit
}

func NewTV() *TV {
	tv := TV{}
	return &tv
}

func (tv *TV) Check(msg *slack.MessageEvent, rtm *slack.RTM) {

	if msg.Text == "tv" {
		rtm.SendMessage(rtm.NewOutgoingMessage("Sended tv event", msg.Channel))
	}
}
