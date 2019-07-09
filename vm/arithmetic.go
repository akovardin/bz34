package vm

type Addition struct {
	vm *VirtualMachine
}

func (cmd Addition) Execute() {
	cmd.vm.Save()
	cmd.vm.RY = cmd.vm.RY + cmd.vm.RX
	cmd.vm.Pop()
}

type Subtraction struct {
	vm *VirtualMachine
}

func (cmd Subtraction) Execute() {
	cmd.vm.Save()
	cmd.vm.RY = cmd.vm.RY - cmd.vm.RX
	cmd.vm.Pop()
}

type Multiplication struct {
	vm *VirtualMachine
}

func (cmd Multiplication) Execute() {
	cmd.vm.Save()
	cmd.vm.RY = cmd.vm.RY * cmd.vm.RX
	cmd.vm.Pop()
}

type Division struct {
	vm *VirtualMachine
}

func (cmd Division) Execute() {
	cmd.vm.Save()
	cmd.vm.RY = cmd.vm.RY / cmd.vm.RX
	cmd.vm.Pop()
}
