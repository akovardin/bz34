package vm

import "math"

type fepowx struct{}

func (fepowx) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Pow(math.E, vm.rx)
}

type flgx struct{}

func (flgx) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Log10(vm.rx)
}

type flnx struct{}

func (flnx) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Log(vm.rx)
}

type ftpowx struct{}

func (ftpowx) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = math.Pow(10, vm.rx)
}
