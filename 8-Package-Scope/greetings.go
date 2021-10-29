package main

import "fmt"

// Slice not Array, because we are not specifying the fixed length
var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("hello", n)
}

// score is coming from main.go, because of "package.main".
func showScore() {
	fmt.Println("you scored this many points", score)
}
