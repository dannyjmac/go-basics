package main

import (
	"fmt"
)

// Maps are kind of like objects in Javascript or Dicts in Python

func main() {

	// Define a map with the key being string and value being floats
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Has to set the value to the type stated in the declaration (float64)
	menu["salad"] = 2.99

}
