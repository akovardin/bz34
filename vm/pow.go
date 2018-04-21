package vm

import "math"

type fsqrt struct{}

func (fsqrt) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Sqrt(vm.rx)
}

type fxpow2 struct{}

func (fxpow2) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Pow(vm.rx, 2)
}

type f1x struct{}

func (f1x) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = 1 / vm.rx
}

type fxpowy struct{}

func (fxpowy) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Pow(vm.rx, vm.ry)
}
