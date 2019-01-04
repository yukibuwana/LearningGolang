package main

func main() {
	cards := newDeck()
	cards.saveToFile("my-cards")

	myCard := newDeckFromFile("my-cards")
	myCard.print()
}
