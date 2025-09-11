package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Trifoi"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen", "Ace"}

	for i := range cardValues {
		for j := range cardSuits {
			cards = append(cards, cardValues[i]+" of "+cardSuits[j])
		}
	}

	return cards
}

func (d deck) print() {
	for i := range d {
		println(d[i])
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) deal(handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) saveToFile(fileName string) {
	err := os.WriteFile(fileName, []byte(d.toString()), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	sDeck := strings.Split(string(byteSlice), ",")

	return deck(sDeck)
}

func (d deck) shuffle() {
	source := rand.NewSource(int64(time.Now().UnixNano()))
	rand := rand.New(source)

	for i := range d {
		newPosition := rand.Intn(len(d))

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) similar () int {
	other := newDeckFromFile("cards")

	if len(d) != len(other) {
		return 0
	}

	count := 0
	for i := range d {
		if d[i] == other[i] {
			count++
		}
	}

	return count
}