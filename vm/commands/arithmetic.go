package commands

import "github.com/horechek/bz34/vm"

type Addition struct{}

func (crx Addition) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RY = vm.RY + vm.RX
	vm.Pop()
}

type Subtraction struct{}

func (crx Subtraction) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RY = vm.RY - vm.RX
	vm.Pop()
}

type Multiplication struct{}

func (crx Multiplication) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RY = vm.RY * vm.RX
	vm.Pop()
}

type Division struct{}

func (crx Division) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RY = vm.RY / vm.RX
	vm.Pop()
}
