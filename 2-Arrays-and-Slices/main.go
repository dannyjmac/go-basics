package main

import "fmt"

func main() {

	// ARRAYS

	// Below is a varialble declaration with type arrray of 3 items all of int
	// You have to have the types on both sides - weird he knows, but you don't have to do it on the left, it's inferred
	// Notice the array has curcly brackets not square
	// An Array has to have a fixed length, you can't change it later
	var ages [3]int = [3]int{20, 25, 30}

	// Notice - not allowed
	// var ages [3]int = [3]int{20, 25, 30, 45}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// SLICES
	// Slices use arrays under the hood, buyt make it easy for you
	// If you do the below (without sopecifying a number) - it auto creates an slice
	var scores = []int{100, 50, 60}
	scores[2] = 25

	// To append an item you have to assign the variable again, because it does not overwrite
	// Because it has to create a new slice
	scores = append(scores, 85)

	// Get a range of a previously declared slice
	// Second the 4th item
	rangeOne := names[1:3]

	// 3rd to last item
	rangeTwo := names[2:]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
}
