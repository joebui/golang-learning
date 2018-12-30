package main

// import "fmt"

func main() {
	cards := newDesk()
	// hand, remainingCard := deal(cards, 5)
	// hand.print()
	// fmt.Println("----")
	// remainingCard.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("test.txt")

	// newCards := newDeskFromFile("test.txt")
	// newCards.print()
	cards.shuffle()
	cards.print()
}
