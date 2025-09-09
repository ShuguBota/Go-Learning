package main

func main() {
	similarCards := 0

	for similarCards < 10 {
		cards := newDeckFromFile("cards")
		cards.shuffle()
		// cards.print()

		similarCards = cards.similar()
		print("There are ", similarCards, " similar cards in the deck\n")
		
		cards.saveToFile("cards")
	}	
}	