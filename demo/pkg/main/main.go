package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("Hello from display")
	msg.Exciting("an exciting message")
}
