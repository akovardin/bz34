package commands

import (
	"math"

	"github.com/horechek/bz34/vm"
)

type Fsqrt struct{}

func (Fsqrt) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Sqrt(vm.RX)
}

type Fxpow2 struct{}

func (Fxpow2) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Pow(vm.RX, 2)
}

type F1x struct{}

func (F1x) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = 1 / vm.RX
}

type Fxpowy struct{}

func (Fxpowy) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Pow(vm.RX, vm.RY)
}
