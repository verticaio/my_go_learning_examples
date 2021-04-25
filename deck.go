package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create new type of deck  which is slice of strings

type deck []string

// create new Deck and fill it
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Tests"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}
	return  cards
}

// get deck content, it is  func with receiver
func ( d deck) print() {
	for i, card := range d {
        fmt.Println(i , card)
	}
}

// It is func with get arguments and return two value
func deal(d deck , handSize int) (deck , deck){
	return d[:handSize], d[handSize:]
}

// deck --> []string --> string --> []byte
func  (d deck) toString() string {
	return strings.Join([]string(d),",")

}

//string --> []byte --WriteToFile
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read from file , convert bye to string to split strings and convert to deck slice type
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// opt 1  - log the error and return a call tto newDeck
		// opt 2 -  log the error and quit program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",") //	Ace of Spades, Two of Spades and etc
	return deck(s)
}

// mix cards in deck
func (d deck) shuffle()  {
	for i := range d {
		// generate sample number between 0 and lenght of slice (-1)
		newPosition := rand.Intn(len(d) - 1)
		// change index number data with generated index number data
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	
}
