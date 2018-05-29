package main

import "fmt"

type bot interface {
	getGreeting() string
}

//declaration of type struct (no fields specified)
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//not gonna have return value -- just receive the arugment and then print something out
//func printGreeting(eb englishBot) {
//pass in the english both, call getGreeting which returns the language greeting in that particular language and print in commandline
//refactor - use only one Println to call either eb or sb
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb spanishBot) {
//pass in the spanish bot, call getGreeting which returns the language greeting in that particular language and print in commandline
//  fmt.Println(sb.getGreeting())
//}

//both types have functions that work on different types
//receiver for func getGreeting
//defined a variable eb, sb without using the value, optional change -- if not making use of actual value, just use the receive type
func (englishBot) getGreeting() string {
	//very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
