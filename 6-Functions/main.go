package main

import (
	"fmt"
)

// We can only pass a string
func sayGreeting(n string) {
	fmt.Printf("Good morning %v", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v", n)
}

// must take in both a string and function with a string parameter
func cycleNames(n []string, f func(string)) {

}

func main() {
	// Bog standard like JS
	sayGreeting("Danny")
	sayBye("Danny")
}
