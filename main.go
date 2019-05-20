package main

import "fmt"

func main() {
	// one way to create a map
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// empty map
	// var colors map[string]string

	// using 'make'
	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// delete "white" from the colors map
	// delete(colors, "white")
	// fmt.Println(colors)

	// iterating through a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
