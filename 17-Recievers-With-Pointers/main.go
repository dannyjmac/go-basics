package main

import "fmt"

func main() {
	mybill := newBill("Marios bill")

	// These now work to add them to the bill becvause they are using reciever functions with pointers
	mybill.addItem("onion soup", 10)
	mybill.updateTip(11)

	fmt.Println(mybill.format())
}
