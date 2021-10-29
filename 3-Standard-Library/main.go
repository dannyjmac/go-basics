package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// strings package - work with strings

	// Below checks if a string is contained, returns a boolean
	fmt.Println(strings.Contains(greeting, "hello"))
	// This substitues strings, returns a new string
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// Uppercase
	fmt.Println(strings.ToUpper(greeting))
	// Get position of string as number
	fmt.Println(strings.Index(greeting, "ll"))

	// sort package - work with slices

	// With Ints
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages)

	// You will notice this time, it has altered the original!
	fmt.Println(ages)

	// With Strings
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	//  alphabetical order
	sort.Strings(names)
	fmt.Println(names)

	// Search all strings for entry
	fmt.Println(sort.SearchStrings(names, "bowser"))
}
