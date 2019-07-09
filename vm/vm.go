package vm

import (
	"fmt"
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
	vm.Execute()
}

func (vm *VirtualMachine) Execute() {
	instr := vm.Memory[vm.pc]

	// execute_ = instr

	vm.pc++
}
