// https://www.youtube.com/watch?v=LBLN4Wu5U4w

package main

import "fmt"

// Go is a Pass-by-value language
// This means it makes copies of values when passed into functions

func main() {

	name := "tifa"

	updateName(name)

	// When printing this, it does not print the name updated by the updateName method.
	// This is because the name passed to updateName is a COPY and not the variable itself.

	// This is like Javascript, doesn't change the original value stored in memory.
	// Instead it creates a new space in memory

	// If you want to overwrite the name, you would need to do this:
	// name = updateName(name)

	fmt.Println(name)

	// HOWEVER - NOT THE CASE WITH MAPS
	// A Map create a pointer in memory to the original, which means if you pass it to a method
	// and then edit the map, it looks at the pointer, then goes to original in mem
	// and updates the original

}

func updateName(x string) {
	x = "wedge"
}
