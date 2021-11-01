package main

// Like a javascript interface
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills, taking in a name, returning a bill struct
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}
