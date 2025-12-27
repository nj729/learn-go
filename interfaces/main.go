package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (e englishBot) getGreeting() string {
	// custom logic
	return "Hello there"
}

func (s spanishBot) getGreeting() string {
	// custom logic
	return "Hola"
}

func printGreeting(b bot) {
	// custom logic
	fmt.Println(b.getGreeting())
}

// func printGreeting(s spanishBot) {
// 	// custom logic
// 	fmt.Println(s.getGreeting())
// }
