package machine

import (
	itemPkg "anksus/vending-machine-go/pkg/models"
	"errors"
)

type PaymentState struct {
	SelectedItem *itemPkg.Item
}

func (ps *PaymentState) SelectItem(wm *VendingMachine, itemName string) error {
	return errors.New("item already selected")
}

func (ps *PaymentState) InsertMoney(wm *VendingMachine, money int64) error {
	inventory := wm.GetInventory()
	if money >= ps.SelectedItem.Price {
		wm.Balance += money
		inventory[ps.SelectedItem.GetName()].Remove()
		wm.SetState(&Dispense{SelectedItem: ps.SelectedItem})
		return nil
	}
	return errors.New("insufficient money")
}

func (ps *PaymentState) DispenseItem(wm *VendingMachine) error {
	return errors.New("payment not completed")
}
