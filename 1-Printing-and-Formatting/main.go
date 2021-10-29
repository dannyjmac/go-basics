package main

import "fmt"

func main() {

	// Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	// This will be a new line because of \n
	fmt.Print("new line")

	// To automatically get a new line
	fmt.Println("new line")
	fmt.Println("and another new line")

	// Outputting variables
	age := 30
	name := "Danny"
	fmt.Println("my age is", age, "and my name is", name)

	// Formatted strings are better than the above

	fmt.Printf("my age is %v and my name is %v", age, name)

	// There are many specifiers you can use
	// %v - standard variable
	// %q - quotes string
	// %T - gives the type of the variable
	// %f - float
	// %0.1f - round to 1 decimal float

	// SprintF - saves formatted strings
	var str = fmt.Sprintf(" my age is %v and my name is %v", age, name)
	fmt.Println(" the saved string is:", str)
}
