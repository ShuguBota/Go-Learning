package main

import "fmt"

func main() {
	// The lines are equivalent

	// var card string = newCard()
	card := newCard(" bob") // We rely on the compiler to infer the type
	// := you use when you declare a variable (at initialization) then we use only =
	fmt.Println(card)

	
	words := [] string {"abc", "def"}
	words = append(words, "ghi")

	for i := range words {
		println(words[i])
	}	
}

func newCard(a string) string {
	return "Ace of Spades" + a
}