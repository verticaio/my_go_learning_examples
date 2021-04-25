package main

import (
	"fmt"
	"os"
)

//deckSize := 60 // Why I not used in main func ? because it is short declaration
var deckSize int


func main () {
    deckSize = 15
    fmt.Println(deckSize)
    var card string = "Ace of Spades"
	fmt.Println(card)
	// card0 = "card0" It won't work bcs not declared previously
	card1 := "Five of Diamonds"  // It detects type
	fmt.Println(card1)
	card1 = "Six of Diamonds"  // Assing again value and not use := for second usage
	fmt.Println(card1)
	fmt.Println("-----------------------------------------------------------------")
	card1 = "Eight of Diamonds"  // Assing again value and not use := for second usage
	fmt.Println(card1)
	fmt.Println(deckSize)
	fmt.Println("-----------------------------------------------------------------")
	card2 := newCard() // call other function in same go file
	fmt.Println(card2)

	fmt.Println("-----------------------------------------------------------------")
	cards := []string{newCard(), "Six of Diamonds", "Seven of Diamonds"}   // [] define slice , string - define content types , curly braces - show content of slice, I call func instead of write myself
	cards = append(cards, "New Card1 Type")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println("-----------------------------------------------------------------")
	b_cards := deck{newCard(), "Card01", "Card02"}
	b_cards = append(b_cards, "New Card3 Type")
	for i, card := range b_cards {
		fmt.Println(i, card)
	}
	fmt.Println("-----------------------------------------------------------------")
	c_cards := deck{newCard(), "Card10", "Card11"}
	c_cards = append(c_cards, "New Card3 Type")
	c_cards.print()

	fmt.Println("\n Create sum of cards in slice   and Print all cards which newly created")
	fmt.Println("-----------------------------------------------------------------")
	d_cards := newDeck()
	fmt.Println(d_cards)
	d_cards.print()


	fmt.Println("\n Share cards between 2 people")
	fmt.Println("-----------------------------------------------------------------")
	//Why I use as  d_cards.deal(3) ? yes I understood I get as argument in that function not receiver why ?
	hand , remaningCards := deal(d_cards, 3)  //data ni split edir ve get edib value lere verir, daha sonra bu value larla self le print metodunu cagirib chapa verir
	hand.print()
	fmt.Println("----------------")
	remaningCards.print()

	fmt.Println("\n write cards to file ")
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println(d_cards.toString())
	d_cards.saveToFile("cards_converted_byte_format")

	fmt.Println("\n read cards from file ")
	cards_from_file := newDeckFromFile("cards_converted_byte_format")
	fmt.Println(cards_from_file)
	cards_from_file.print()

	fmt.Println("\n Mix the cards in deck ")
	cards_from_file.shuffle()
	cards_from_file.print()
	fmt.Println(len(d_cards))
	os.Remove("cards_converted_byte_format")
}

func newCard() string {
	return "Ten of Diamonds"
}