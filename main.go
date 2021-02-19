//Executable file "Main"
package main

func main() {

	//Colon only used when initializing variable
	//Variable with operator that relies on interpretter to identify the data type based on value being assigned
	//variable is being assign to a slice of type string with values of
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// initializing cards variable and assigning value returned from newDeck() function
	// cards := newDeck()
	//passing deck type of string cards to receiver of saveToFile to write to file
	//passing filename string into function to specify the filename to write
	// cards.savfeToFile("my_cards")

	//new deck of cards
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
