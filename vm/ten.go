package vm

import (
	"math"
)

type Fepowx struct {
	vm *VirtualMachine
}

func (cmd Fepowx) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Pow(math.E, cmd.vm.RX)
}

type Flgx struct {
	vm *VirtualMachine
}

func (cmd Flgx) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Log10(cmd.vm.RX)
}

type Flnx struct {
	vm *VirtualMachine
}

func (cmd Flnx) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Log(cmd.vm.RX)
}

type Ftpowx struct {
	vm *VirtualMachine
}

func (cmd Ftpowx) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Pow(10, cmd.vm.RX)
}
