package commands

import (
	"math"

	"github.com/horechek/bz34/vm"
)

type Fepowx struct{}

func (Fepowx) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Pow(math.E, vm.RX)
}

type Flgx struct{}

func (Flgx) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Log10(vm.RX)
}

type Flnx struct{}

func (Flnx) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Log(vm.RX)
}

type Ftpowx struct{}

func (Ftpowx) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Pow(10, vm.RX)
}
