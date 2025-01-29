package main

import (
	"anksus/vending-machine-go/pkg/machine"
	"anksus/vending-machine-go/pkg/models"
	"fmt"
)

func main() {
	wm := machine.NewVendingMachine()
	coke := models.Item{
		Count: 2,
		Name:  "Coke",
		Price: 20,
	}
	redBull := models.Item{
		Count: 1,
		Name:  "RedBull",
		Price: 100,
	}
	wm.AddItem(coke)
	wm.AddItem(redBull)

	fmt.Println(wm.GetItems())
	err := wm.SelectItem("Coke")

	if err != nil {
		fmt.Println(err)
	}
	err = wm.InsertMoney(30)

	if err != nil {
		fmt.Println(err)
	}
	err = wm.DispenseItem()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("printing vm %+v\n", wm.Items["Coke"])
}
