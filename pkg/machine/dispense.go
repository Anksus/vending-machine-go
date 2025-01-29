package machine

import (
	itemPkg "anksus/vending-machine-go/pkg/models"
	"errors"
	"fmt"
)

type Dispense struct {
	SelectedItem *itemPkg.Item
}

func (ds *Dispense) SelectItem(wm *VendingMachine, itemName string) error {
	return errors.New("item already selected")
}

func (ds *Dispense) InsertMoney(wm *VendingMachine, money int64) error {
	return errors.New("payment already completed")
}

func (ds *Dispense) DispenseItem(wm *VendingMachine) error {
	fmt.Println("Dispensing item: ", ds.SelectedItem.GetName())
	wm.SetState(&IdleState{})
	return nil
}
