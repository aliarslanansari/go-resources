package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades" //exactly same as above code | Data type of card will be set to string as we are assigning a string to it. (:=)
	card = newCardString()  // we can assign only string to it as during declaration we have set it to string | newCardString() func returns a string
	newCard := newCardInt() // data type of newCard variable is automatically set to int as newCardInt() func returns int.
	fmt.Println(card)
	fmt.Println(newCard)
}

func newCardString() string {
	return "Five of Diamonds"
}

func newCardInt() int {
	return 12
}
