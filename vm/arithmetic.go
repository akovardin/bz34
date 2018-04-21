package vm

type addition struct{}

func (crx addition) execute(vm *VirtualMachine) {
	vm.save()
	vm.ry = vm.ry + vm.rx
	vm.pop()
}

type subtraction struct{}

func (crx subtraction) execute(vm *VirtualMachine) {
	vm.save()
	vm.ry = vm.ry - vm.rx
	vm.pop()
}

type multiplication struct{}

func (crx multiplication) execute(vm *VirtualMachine) {
	vm.save()
	vm.ry = vm.ry * vm.rx
	vm.pop()
}

type division struct{}

func (crx division) execute(vm *VirtualMachine) {
	vm.save()
	vm.ry = vm.ry / vm.rx
	vm.pop()
}
