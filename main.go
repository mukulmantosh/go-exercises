package main

func main() {
	//var number = mobile{}
	//fmt.Println(number.mobileNumber("+91"))

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//err := cards.saveToFile("my_cards.txt")
	//if err != nil {
	//	return
	//}

	//readFile := cards.newDeckFromFile("my_cards.txt")
	//fmt.Println(readFile)

	//cards.print()
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//fmt.Println("    ")
	//remainingCards.print()
	//fmt.Println(cards.toString())

	//greeting := "Hi there"
	//fmt.Println([]byte(greeting))

}
