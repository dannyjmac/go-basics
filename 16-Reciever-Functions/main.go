package main

import "fmt"

func main() {
	mybill := newBill("Marios bill")

	// We want methods for the structs like so: mybill.format()

	// Now that a reciever function is written in bill.go, I cannot do this format() , only mybill.format()
	fmt.Println(mybill.format())
}
