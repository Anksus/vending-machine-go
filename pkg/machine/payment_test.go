package machine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectItemPayment(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.SelectItem(vm, "Coke")
	err = vm.State.SelectItem(vm, "Coke")

	assert.EqualError(t, err, "item already selected")

}

func TestInsertMoneyPaymentSuccess(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.SelectItem(vm, "Coke")
	err = vm.State.InsertMoney(vm, 123)

	assert.NoError(t, err)
}

func TestInsertMoneyPaymentFailure(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.SelectItem(vm, "Coke")
	err = vm.State.InsertMoney(vm, 1)

	assert.EqualError(t, err, "insufficient money")
}

func TestDispenseItemPayment(t *testing.T) {
	vm := getVMWithItem()

	err := vm.State.SelectItem(vm, "Coke")
	err = vm.State.DispenseItem(vm)
	assert.EqualError(t, err, "payment not completed")
}
