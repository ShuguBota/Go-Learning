package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	englishBot := englishBot{}
	spanishBot := spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "¡Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

type bot interface {
	getGreeting() string // The types that want to use this type then need to implement this method
	// The interface is satisfied implicitly (no need to declare it explicitly "implements bot")
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}