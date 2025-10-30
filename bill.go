package main

import "fmt"

type bill struct {
	names string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	return bill{
		names: name,
		items: map[string]float64{
			"soup": 4.5,
			"milk": 8.95,
		},
		tip: 0,
	}

}
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ....$%f\n", "total:", total+b.tip)
	return fs

}

// updateTip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(key string, price float64) {
	b.items[key] = price
}

// updateName
func updateName(name *string) {
	*name = "laudia"

}
