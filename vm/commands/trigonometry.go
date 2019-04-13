package commands

import (
	"math"

	"github.com/horechek/bz34/vm"
)

type Farcsin struct{}

func (Farcsin) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Asin(vm.RX))
}

type Farccos struct{}

func (Farccos) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Acos(vm.RX))
}

type Farctg struct{}

func (Farctg) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Atan(vm.RX))
}

type Fpi struct{}

func (Fpi) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = math.Pi
}

type Fsin struct{}

func (Fsin) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Sin(vm.RX))
}

type Fcos struct{}

func (Fcos) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Cos(vm.RX))
}

type Ftg struct{}

func (Ftg) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.Transform(math.Tan(vm.RX))
}
