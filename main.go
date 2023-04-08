package main

func main() {
	// Assign variable of type string
	// var card string = "Ace of Spades"
	// Shorthand:

	cards := newDeck()
	cards.shuffle()
	
	cards.print()
	// cards.saveToFile("cards.csv")
}