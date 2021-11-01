// // https://www.youtube.com/watch?v=4B2rwYvuiBo

package main

import "fmt"

func main() {
	name := "tifa"
	// You can get the moery location of a variable by using an ampersand
	fmt.Println("Memory qaddress of name is: ", &name)
	// This can be stored into a variable, which creates its own memory address, holding another one
	m := &name

	//
	fmt.Println("Value at memory address instead of just the memory address", *m)

	// Now if you want to update the original stored value in memory
	// If you passed just name, then it would not update, only the in scope copy
	updateName(m)
	// This has now updated the original
	fmt.Println(name)
}

// When there is an asterisk infront of the type and variable it means to type for / get the original value in the pointer
func updateName(x *string) {
	*x = "wedge"
}
