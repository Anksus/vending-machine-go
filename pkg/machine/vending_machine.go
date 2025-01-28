package machine

import (
	"anksus/vending-machine-go/pkg/item"
	"fmt"
)

type VendingMachine struct {
	Items   []item.Item
	Balance *Balance
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		Items:   make([]item.Item, 0),
		Balance: NewBalance(),
	}
}

func (v *VendingMachine) AddItem(i item.Item) {
	v.Items = append(v.Items, i)
}

func (v *VendingMachine) GetItems() []string {
	var items []string
	for _, i := range v.Items {
		items = append(items, i.GetName())
	}
	return items
}

func (v *VendingMachine) DispenseItem(name string, balance *Balance) {
	for _, v := range v.Items {
		if v.GetName() == name {
			err := v.Remove()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Dispensing item: ", name)
			return
		}
	}
	fmt.Println("Item not found")
}
