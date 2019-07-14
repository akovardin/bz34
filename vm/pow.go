package vm

import (
	"math"
)

type Fsqrt struct {
	vm *VirtualMachine
}

func (cmd Fsqrt) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Sqrt(cmd.vm.RX)
}

type Fxpow2 struct {
	vm *VirtualMachine
}

func (cmd Fxpow2) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Pow(cmd.vm.RX, 2)
}

type F1x struct {
	vm *VirtualMachine
}

func (cmd F1x) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = 1 / cmd.vm.RX
}

type Fxpowy struct {
	vm *VirtualMachine
}

func (cmd Fxpowy) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Pow(cmd.vm.RX, cmd.vm.RY)
}
