package main

import (
	"fmt"

	"github.com/mogeta/irkit/relay"
	"github.com/mogeta/leap"
)

var irkit *relay.Irkit
var leapMotion *leap.LeapMotion

type leapToIrkit struct {
	leap.Action
}

func main() {
	irkit = relay.New()

	lti := &leapToIrkit{}

	leapMotion = leap.New(lti)
	leapMotion.Start()

}

func (l *leapToIrkit) Swipe() {
	var data = `{"format":"raw","freq":38,"data":[17421,8755,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190,1037,1150,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,1037,1150,3228,1150,1037,1190,3228,1190,1037,1190,1037,1190,3228,1150,1037,1190,1037,1190,1037,1190,3228,1190,1037,1190,3228,1190,3228,1190,1037,1190,3228,1150,3228,1150,3228,1150,65535,0,13693,17421,4400,1150]}`
	err := irkit.SendMessage(data)
	if err != nil {
		fmt.Println("hoge")
	}
}

func (l *leapToIrkit) Circle() {

}
