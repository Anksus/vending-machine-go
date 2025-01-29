package machine

type State interface {
	SelectItem(wm *VendingMachine, itemName string) error
	InsertMoney(wm *VendingMachine, money int64) error
	DispenseItem(wm *VendingMachine) error
}
