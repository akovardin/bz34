package vm

import (
	"math"
)

type Farcsin struct {
	vm *VirtualMachine
}

func (cmd Farcsin) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Asin(cmd.vm.RX))
}

type Farccos struct {
	vm *VirtualMachine
}

func (cmd Farccos) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Acos(cmd.vm.RX))
}

type Farctg struct {
	vm *VirtualMachine
}

func (cmd Farctg) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Atan(cmd.vm.RX))
}

type Fpi struct {
	vm *VirtualMachine
}

func (cmd Fpi) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = math.Pi
}

type Fsin struct {
	vm *VirtualMachine
}

func (cmd Fsin) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Sin(cmd.vm.RX))
}

type Fcos struct {
	vm *VirtualMachine
}

func (cmd Fcos) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Cos(cmd.vm.RX))
}

type Ftg struct {
	vm *VirtualMachine
}

func (cmd Ftg) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.Transform(math.Tan(cmd.vm.RX))
}
