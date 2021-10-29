package main

// Uses variables/functions declared in another file (greetings.go), because they are both "package main"

// In order to run this, you need to run both files together so:
// go run main.go greetings.go

// Has to be outside main(), as main() is it's own scope, so has to be in the root.
var score = 99.5

func main() {
	sayHello("mario")

	showScore()
}
