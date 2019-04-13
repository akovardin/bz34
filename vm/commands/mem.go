package commands

import "github.com/horechek/bz34/vm"

type Fo struct{}

func (Fo) Execute(vm *vm.VirtualMachine) {
	vm.Circle()
}

type Fp struct {
	rname string
}

func (f Fp) Execute(vm *vm.VirtualMachine) {
	vm.RR[f.rname] = vm.RX
}

type Fip struct {
	rname string
}

func (f Fip) Execute(vm *vm.VirtualMachine) {
	vm.Save()
	vm.RX = vm.RR[f.rname]
}
