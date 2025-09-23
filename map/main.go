package main

import (
	"fmt"
)

func main() {
	colors := map[string]string {
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff", // Trailing comma is required
	}

	colors["blue"] = "#0000ee" // Update value

	fmt.Println(colors)

	for color, hex := range colors {
		fmt.Println(color, "->", hex)
	}

	// Another way to create a map
	colors2 := make(map[string]string)

	colors2["white"] = "#ffffff"
	colors2["black"] = "#000000"

	for color, hex := range colors2 {
		fmt.Println(color, "->", hex)
	}

	fmt.Println(colors2)

	fmt.Println("Deleting white")
	delete(colors2, "white")

	fmt.Println(colors2)
}