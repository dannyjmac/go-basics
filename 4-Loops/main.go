package main

import (
	"fmt"
)

func main() {
	x := 0

	// Pretty bog standard
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	// Like a js forloop
	for i := 0; i < 5; i++ {
		fmt.Println("vlaue of i is:", i)
	}

	// Similar to python
	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Loop through specific slice, and give the value
	for index, value := range names {
		fmt.Println(index, value)
	}

	// Don't want to use index
	for _, value := range names {
		fmt.Println(value)
	}

}
