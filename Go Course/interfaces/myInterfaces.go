package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
	//otherFunc(int) (string, error)
}

type engBot struct{}
type spaBot struct{}

func main() {
	e := engBot{}
	s := spaBot{}

	printGreeting(e)
	printGreeting(s)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (engBot) getGreeting() string {
	return "Hi there!"
}

func (spaBot) getGreeting() string {
	return "Hola!"
}
