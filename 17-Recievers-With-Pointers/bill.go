package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// This is a reciever function which takes in a pointer to bill struct, not just a bill struct
func (b *bill) updateTip(tip float64) {
	b.tip = tip

	// Which is actually doing this behind the scenes
	// (*b).tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// We are now passing in pointers for all methods, as its more efficient to not copy the new bill struct to another memory slot
// Instead it's just a pointer
func (b *bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	fmt.Println("tip", b.tip)

	return fs
}
