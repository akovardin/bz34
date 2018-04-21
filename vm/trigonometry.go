package vm

import "math"

type farcsin struct{}

func (farcsin) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Asin(vm.rx))
}

type farccos struct{}

func (farccos) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Acos(vm.rx))
}

type farctg struct{}

func (farctg) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Atan(vm.rx))
}

type fpi struct{}

func (fpi) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Pi
}

type fsin struct{}

func (fsin) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Sin(vm.rx))
}

type fcos struct{}

func (fcos) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Cos(vm.rx))
}

type ftg struct{}

func (ftg) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.transform(math.Tan(vm.rx))
}
