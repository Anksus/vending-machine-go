package machine

import (
	itemPkg "anksus/vending-machine-go/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getVMWithItem() *VendingMachine {
	vm := NewVendingMachine()
	item := itemPkg.Item{
		Price: 10,
		Name:  "Coke",
		Count: 1,
	}
	vm.AddItem(item)

	return vm
}

func TestSelectItem(t *testing.T) {
	vm := getVMWithItem()

	err := vm.SelectItem("Coke")
	err = vm.InsertMoney(123)

	err = vm.State.SelectItem(vm, "Coke")

	assert.EqualError(t, err, "item already selected")
}

func TestInsertMoney(t *testing.T) {
	vm := getVMWithItem()

	err := vm.SelectItem("Coke")
	err = vm.InsertMoney(123)

	err = vm.State.InsertMoney(vm, 123)

	assert.EqualError(t, err, "payment already completed")
}

func TestDispenseItem(t *testing.T) {
	vm := getVMWithItem()

	err := vm.SelectItem("Coke")
	err = vm.InsertMoney(123)

	err = vm.State.DispenseItem(vm)

	assert.NoError(t, err)
}
