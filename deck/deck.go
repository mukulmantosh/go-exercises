package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck'
// which is a slice of strings

type deck []string

//type mobile []string

//func (num mobile) mobileNumber(extension string) string {
//	return extension + "1234567890"
//}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	data := []string(d)
	return strings.Join(data, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	result := strings.Split(string(data), ",")
	return deck(result)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := range d {
		newPosition := random.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
