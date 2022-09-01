package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 2)

	hand.print()
	remainingDeck.print()

	cards.saveToFile("myCards")

	fmt.Println(loadFromFile("myCards"))

	cards.print()
	cards.shuffle()
	cards.print()

	//evenOrOdd()

	// createPerson()

	//executeMaps()
}

func evenOrOdd() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var intSlice [10]int

	for i := 0; i < 10; i++ {
		intSlice[i] = r.Intn(100)
	}

	for i := range intSlice {

		if intSlice[i]%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
