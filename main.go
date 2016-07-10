package main

import (
	"github.com/mogeta/irkit/relay"
	"github.com/mogeta/leap"
)

var irkit *relay.Irkit
var leapMotion *leap.LeapMotion

func main() {
	irkit = relay.New()
	leapMotion = leap.New()
}
