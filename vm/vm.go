package vm

import (
	"fmt"
	"log"
	"math"
)

type VirtualMachine struct {
	RX float64 // текущий регистр
	RY float64 // рабочий регистр
	// стековые регистры
	RZ float64
	RT float64

	// регистр предыдущего результата
	RX1 float64

	// регистры общего назначеия
	RR map[string]float64

	// memmory
	Memory [98]int

	Grad bool

	// program counter
	pc int

	debug bool
	run   bool

	commands map[int]Command
}

func NewVirtualMachine() *VirtualMachine {
	v := &VirtualMachine{
		RR: map[string]float64{
			"0": 0,
			"1": 0,
			"2": 0,
			"3": 0,
			"4": 0,
			"5": 0,
			"6": 0,
			"7": 0,
			"8": 0,
			"9": 0,
			"A": 0,
			"B": 0,
			"C": 0,
			"D": 0,
		},
	}

	v.commands = map[int]Command{
		0x0d: Clear{vm: v},
		0x0e: Enter{vm: v},
		0x10: Addition{vm: v},
		0x11: Subtraction{vm: v},
		0x12: Multiplication{vm: v},
		0x13: Division{vm: v},

		0x15: Ftpowx{vm: v},
		0x16: Fepowx{vm: v},
		0x17: Flgx{vm: v},
		0x18: Flnx{vm: v},

		0x19: Farcsin{vm: v},
		0x1a: Farccos{vm: v},
		0x1b: Farctg{vm: v},
		0x1c: Fsin{vm: v},
		0x1d: Fcos{vm: v},
		0x1e: Ftg{vm: v},
		0x20: Fpi{vm: v},

		0x21: Fsqrt{vm: v},
		0x22: Fxpow2{vm: v},
		0x23: F1x{vm: v},
		0x24: Fxpowy{vm: v},
		0x25: Fo{vm: v},

		0x40: Fp{v, "0"},
		0x41: Fp{v, "1"},
		0x42: Fp{v, "2"},
		0x43: Fp{v, "3"},
		0x44: Fp{v, "4"},
		0x45: Fp{v, "5"},
		0x46: Fp{v, "6"},
		0x47: Fp{v, "7"},
		0x48: Fp{v, "8"},
		0x49: Fp{v, "9"},
		0x4a: Fp{v, "A"},
		0x4b: Fp{v, "B"},
		0x4c: Fp{v, "C"},
		0x4d: Fp{v, "D"},

		0x60: Fip{v, "0"},
		0x61: Fip{v, "1"},
		0x62: Fip{v, "2"},
		0x63: Fip{v, "3"},
		0x64: Fip{v, "4"},
		0x65: Fip{v, "5"},
		0x66: Fip{v, "6"},
		0x67: Fip{v, "7"},
		0x68: Fip{v, "8"},
		0x69: Fip{v, "9"},
		0x6a: Fip{v, "A"},
		0x6b: Fip{v, "B"},
		0x6c: Fip{v, "C"},
		0x6d: Fip{v, "D"},

		0x70: Stop{v},
	}

	return v
}

// результаты всех вычислений храняться в регистре X
func (vm *VirtualMachine) Result() float64 {
	return vm.RX
}

func (vm *VirtualMachine) Reset() {
	vm.Grad = false
	vm.debug = false

	vm.RX = 0
	vm.RY = 0
	vm.RZ = 0
	vm.RT = 0
	vm.RX1 = 0

	for name := range vm.RR {
		vm.RR[name] = 0
	}

	vm.Memory = [98]int{}
}

func (vm *VirtualMachine) Dump() {
	fmt.Println("x =", vm.RX)
	fmt.Println("y =", vm.RY)
	fmt.Println("z =", vm.RZ)
	fmt.Println("t =", vm.RT)
	fmt.Println()

	fmt.Println("x1 =", vm.RX1)
	fmt.Println()

	fmt.Println("0 =", vm.RR["0"])
	fmt.Println("1 =", vm.RR["1"])
	fmt.Println("2 =", vm.RR["2"])
	fmt.Println("3 =", vm.RR["3"])
	fmt.Println("4 =", vm.RR["4"])
	fmt.Println("5 =", vm.RR["5"])
	fmt.Println("6 =", vm.RR["6"])
	fmt.Println("7 =", vm.RR["7"])
	fmt.Println("8 =", vm.RR["8"])
	fmt.Println("9 =", vm.RR["9"])
	fmt.Println("A =", vm.RR["A"])
	fmt.Println("B =", vm.RR["B"])
	fmt.Println("C =", vm.RR["C"])
	fmt.Println("D =", vm.RR["D"])
}

func (vm *VirtualMachine) Transform(val float64) float64 {
	if vm.Grad {
		val = val / (math.Pi / 180)
	}

	return val
}

// функции для работы с стековыми и рабочими регистрами
func (vm *VirtualMachine) Save() {
	vm.RX1 = vm.RX
}

func (vm *VirtualMachine) Bx() {
	vm.Put()
	vm.RX = vm.RX1
}

func (vm *VirtualMachine) Put() {
	vm.RT = vm.RZ
	vm.RZ = vm.RY
	vm.RY = vm.RX
}

func (vm *VirtualMachine) Pop() {
	vm.RX = vm.RY
	vm.RY = vm.RZ
	vm.RZ = vm.RT
}

func (vm *VirtualMachine) Circle() {
	x := vm.RX

	vm.RX = vm.RY
	vm.RY = vm.RZ
	vm.RZ = vm.RT

	vm.RT = x
}

func (vm *VirtualMachine) Run() {
	for {
		vm.Step()
		if !vm.run {
			return
		}
	}
}

func (vm *VirtualMachine) Step() {
	vm.run = true
	if err := vm.Execute(); err != nil {
		log.Println("error on execute", err)
	}
}

func (vm *VirtualMachine) Execute() error {
	instr := vm.Memory[vm.pc]

	cmd, ok := vm.commands[instr]
	if !ok {
		vm.Save()
		vm.Put()
		vm.RX = float64(instr)
		vm.pc++
		return nil
	}

	cmd.Execute()

	vm.pc++
	return nil
}
