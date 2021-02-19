package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new deck type
//that is a slice of strings

//defines new type that extends existing type
type deck []string

//Creates a deck of cards with type deck: slice of strings representing cards
func newDeck() deck {
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}

	//Loop over suits and values to create each card and append to cards slice
	//Use _ to indicate a var that we do not wish to use in our logic: loops
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	//returns the deck type of string -> cards
	return cards
}

//Add new print method to the type deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//A function that requires two args: d deck and handSize int that returns two deck types of slices of deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//toString function that has a reciever of d deck that converts a deck type of strings to one string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//function that writes a deck to a file using the ioutil package with WriteFile function
//Uses the toString() function to convert deck to string
//Then converts string to byte slice
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Read deck from a file. Filename string as required argument to specify the file to read from.
func readDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

//Function that randomly swaps each card with another. I.e. Shuffle cards
func (d deck) shuffle() {
	//time.Now().unixNano() creates a new Int64 number everytime programme starts -> Used as Seed
	//Seed used to create new source object
	//Source object used as basis for New random number generator
	//pseudo random generation -> Research
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
