package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	fmt.Print("Create a new bill name: ")
	// Get everything up untill they press enter
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}
func createBill() bill {
	// Allows users to add information through the terminal
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name", reader)
	fmt.Println("Created the bill", name)

	bill := newBill(name)

	return bill
}

func main() {
	mybill := createBill()
	fmt.Println(mybill)
}
