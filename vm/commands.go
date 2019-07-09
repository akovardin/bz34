package vm

type Command interface {
	Execute()
}

type Clear struct {
	vm *VirtualMachine
}

func (cmd Clear) Execute() {
	cmd.vm.RX = 0
}

type Enter struct {
	vm *VirtualMachine
}

func (cmd Enter) Execute() {
	// сохраняем прошлый результат
	cmd.vm.Save()
	cmd.vm.Put()
}

type Stop struct {
	vm *VirtualMachine
}

func (cmd Stop) Execute() {
	cmd.vm.run = false
}
