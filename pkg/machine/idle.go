package machine

import (
	"fmt"
)

type IdleState struct{}

func (i *IdleState) SelectItem(wm *VendingMachine, itemName string) error {
	inventory := wm.GetInventory()
	if item, ok := inventory[itemName]; ok && item.Count > 0 {
		fmt.Println("Item selected: ", itemName)

		wm.SetState(&PaymentState{SelectedItem: item})
		return nil
	}
	return fmt.Errorf("item %s is not available", itemName)
}

func (i *IdleState) InsertMoney(wm *VendingMachine, money int64) error {
	return fmt.Errorf("no item selected")
}

func (i *IdleState) DispenseItem(wm *VendingMachine) error {
	return fmt.Errorf("no item selected")
}
