package vm

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type VirtualMachine struct {
	rx float64 // текущий регистр
	ry float64 // рабочий регистр
	// стековые регистры
	rz float64
	rt float64

	// регистр предыдущего результата
	rx1 float64

	// регистры общего назначеия
	rr map[string]float64

	grad  bool
	debug bool
}

func NewVirtualMachine() *VirtualMachine {
	v := &VirtualMachine{
		rr: map[string]float64{
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

// преодразование текстового значения в код команды
func (vm *VirtualMachine) Parse(input string) (int, error) {
	val, err := strconv.ParseFloat(input, 64)
	if err == nil {
		vm.save()
		vm.put()
		vm.rx = val
		return 0, nil // enter code
	}

	if input == "dump" {
		vm.dump()
		return 0, nil
	}

	code, ok := tokens[input]
	if !ok {
		return 0, errors.New("undefined token")
	}

	return code, nil
}

// выполняем команду, приводим виртуальную машину к определенному состоянию
func (vm *VirtualMachine) Execute(code int) error {
	if code == 0 {
		return nil
	}

	cmd, ok := comands[code]
	if !ok {
		return errors.New("undefined code")
	}

	cmd.execute(vm)

	return nil
}

// результаты всех вычислений храняться в регистре X
func (vm *VirtualMachine) Result() float64 {
	return vm.rx
}

func (vm *VirtualMachine) Clear() {
	vm.grad = false
	vm.debug = false

	vm.rx = 0
	vm.ry = 0
	vm.rz = 0
	vm.rt = 0
	vm.rx1 = 0

	for name := range vm.rr {
		vm.rr[name] = 0
	}
}

func (vm *VirtualMachine) dump() {
	fmt.Println("x =", vm.rx)
	fmt.Println("y =", vm.ry)
	fmt.Println("z =", vm.rz)
	fmt.Println("t =", vm.rt)
	fmt.Println()

	fmt.Println("x1 =", vm.rx1)
	fmt.Println()

	fmt.Println("0 =", vm.rr["0"])
	fmt.Println("1 =", vm.rr["1"])
	fmt.Println("2 =", vm.rr["2"])
	fmt.Println("3 =", vm.rr["3"])
	fmt.Println("4 =", vm.rr["4"])
	fmt.Println("5 =", vm.rr["5"])
	fmt.Println("6 =", vm.rr["6"])
	fmt.Println("7 =", vm.rr["7"])
	fmt.Println("8 =", vm.rr["8"])
	fmt.Println("9 =", vm.rr["9"])
	fmt.Println("A =", vm.rr["A"])
	fmt.Println("B =", vm.rr["B"])
	fmt.Println("C =", vm.rr["C"])
	fmt.Println("D =", vm.rr["D"])
}

func (vm *VirtualMachine) transform(val float64) float64 {
	if vm.grad {
		val = val / (math.Pi / 180)
	}

	return val
}

// функции для работы с стековыми и рабочими регистрами
func (vm *VirtualMachine) save() {
	vm.rx1 = vm.rx
}

func (vm *VirtualMachine) bx() {
	vm.put()
	vm.rx = vm.rx1
}

func (vm *VirtualMachine) put() {
	vm.rt = vm.rz
	vm.rz = vm.ry
	vm.ry = vm.rx
}

func (vm *VirtualMachine) pop() {
	vm.rx = vm.ry
	vm.ry = vm.rz
	vm.rz = vm.rt
}

func (vm *VirtualMachine) circle() {
	x := vm.rx

	vm.rx = vm.ry
	vm.ry = vm.rz
	vm.rz = vm.rt

	vm.rt = x
}
