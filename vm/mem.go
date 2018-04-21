package vm

type fo struct{}

func (fo) execute(vm *VirtualMachine) {
	vm.circle()
}

type fp struct {
	rname string
}

func (f fp) execute(vm *VirtualMachine) {
	vm.rr[f.rname] = vm.rx
}

type fip struct {
	rname string
}

func (f fip) execute(vm *VirtualMachine) {
	vm.save()
	vm.rx = vm.rr[f.rname]
}
