package main

import (
	"anksus/vending-machine-go/pkg/item"
	"anksus/vending-machine-go/pkg/machine"
)

func main() {
	wm := machine.NewVendingMachine()
	coke := item.Item{
		Count: 10,
		Name:  "Coke",
		Price: 20,
	}
	redBull := item.Item{
		Count: 10,
		Name:  "RedBull",
		Price: 100,
	}
	wm.AddItem(coke)
	wm.AddItem(redBull)

	wm.DispenseItem("Coke")
}
