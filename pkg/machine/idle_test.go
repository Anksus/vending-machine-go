package machine

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectItemIdleSuccessful(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.SelectItem(vm, "Coke")

	assert.NoError(t, err)
}

func TestSelectItemIdleWithWrongItem(t *testing.T) {
	vm := getVMWithItem()
	rb := "RedBull"
	err := vm.State.SelectItem(vm, rb)

	assert.EqualError(t, err, fmt.Sprintf("item %s is not available", rb))
}

func TestInsertMoneyIdle(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.InsertMoney(vm, 123)

	assert.EqualError(t, err, "no item selected")
}

func TestDispenseItemIdle(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.DispenseItem(vm)

	assert.EqualError(t, err, "no item selected")
}
