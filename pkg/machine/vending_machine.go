package machine

import (
	itemPkg "anksus/vending-machine-go/pkg/models"
)

type VendingMachine struct {
	Items   map[string]*itemPkg.Item
	Balance int64
	State   State
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		Items:   make(map[string]*itemPkg.Item),
		Balance: 0,
		State:   &IdleState{},
	}
}

func (v *VendingMachine) AddItem(i itemPkg.Item) {
	v.Items[i.GetName()] = &i
}

func (v *VendingMachine) GetItems() []string {
	var items []string
	for _, i := range v.Items {
		items = append(items, i.GetName())
	}
	return items
}

func (v *VendingMachine) SelectItem(itemName string) error {
	return v.State.SelectItem(v, itemName)
}

func (v *VendingMachine) InsertMoney(money int64) error {
	return v.State.InsertMoney(v, money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.State.DispenseItem(v)
}

func (v *VendingMachine) SetState(s State) {
	v.State = s
}
func (v *VendingMachine) GetInventory() map[string]*itemPkg.Item {
	return v.Items
}
