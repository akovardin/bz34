package vm

type Bp struct {
	vm *VirtualMachine
}

func (cmd Bp) Execute() {
	address := cmd.vm.Memory[cmd.vm.pc+1]
	cmd.vm.pc = address
}
