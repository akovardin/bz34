package vm

type Fo struct {
	vm *VirtualMachine
}

func (cmd Fo) Execute() {
	cmd.vm.Circle()
}

type Fp struct {
	vm    *VirtualMachine
	rname string
}

func (cmd Fp) Execute() {
	cmd.vm.RR[cmd.rname] = cmd.vm.RX
}

type Fip struct {
	vm    *VirtualMachine
	rname string
}

func (cmd Fip) Execute() {
	cmd.vm.Save()
	cmd.vm.RX = cmd.vm.RR[cmd.rname]
}
