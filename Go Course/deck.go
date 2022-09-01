package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

var separator string = ","

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	return deck{"Ace of Spades", "Four of Hearts", "King of Diamonds"}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return strings.Split(string(bs), separator)
}

func (d deck) toString() string {
	return strings.Join([]string(d), separator)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d))

		d[newPosition], d[i] = d[i], d[newPosition]
	}
}
