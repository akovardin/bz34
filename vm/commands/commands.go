package commands

import "github.com/horechek/bz34/vm"

type Command interface {
	Execute(vm *vm.VirtualMachine)
}

var Commands = map[int]Command{
	0x0d: Clear{},
	0x0e: Enter{},
	0x10: Addition{},
	0x11: Subtraction{},
	0x12: Multiplication{},
	0x13: Division{},

	0x15: Ftpowx{},
	0x16: Fepowx{},
	0x17: Flgx{},
	0x18: Flnx{},

	0x19: Farcsin{},
	0x1a: Farccos{},
	0x1b: Farctg{},
	0x1c: Fsin{},
	0x1d: Fcos{},
	0x1e: Ftg{},
	0x20: Fpi{},

	0x21: Fsqrt{},
	0x22: Fxpow2{},
	0x23: F1x{},
	0x24: Fxpowy{},

	0x25: Fo{},
	0x40: Fp{"0"},
	0x41: Fp{"1"},
	0x42: Fp{"2"},
	0x43: Fp{"3"},
	0x44: Fp{"4"},
	0x45: Fp{"5"},
	0x46: Fp{"6"},
	0x47: Fp{"7"},
	0x48: Fp{"8"},
	0x49: Fp{"9"},
	0x4a: Fp{"A"},
	0x4b: Fp{"B"},
	0x4c: Fp{"C"},
	0x4d: Fp{"D"},

	0x60: Fip{"0"},
	0x61: Fip{"1"},
	0x62: Fip{"2"},
	0x63: Fip{"3"},
	0x64: Fip{"4"},
	0x65: Fip{"5"},
	0x66: Fip{"6"},
	0x67: Fip{"7"},
	0x68: Fip{"8"},
	0x69: Fip{"9"},
	0x6a: Fip{"A"},
	0x6b: Fip{"B"},
	0x6c: Fip{"C"},
	0x6d: Fip{"D"},
}

type Clear struct{}

func (crx Clear) Execute(vm *vm.VirtualMachine) {
	vm.RX = 0
}

type Enter struct{}

func (e Enter) Execute(vm *vm.VirtualMachine) {
	// сохраняем прошлый результат
	vm.Save()
	vm.Put()
}
