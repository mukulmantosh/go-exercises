package main

import "fmt"

func main() {
	colors := map[string]string{}
	colors["red"] = "#ff0000"
	colors["green"] = "#4bf745"
	colors["white"] = "#ffffff"

	// Another way of creating map
	//colors := make(map[string]string)
	//colors["red"] = "#ff0000"

	// Delete Key
	//delete(colors, "red")

	printMap(colors)
	fmt.Println(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
