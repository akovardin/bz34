package main

import (
	"fmt"
	"log"

	"github.com/horechek/bz34/vm"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	machine := vm.NewVirtualMachine()

	machine.Memory[0] = 1
	machine.Memory[1] = 10
	machine.Memory[2] = 0x10
	machine.Memory[3] = 20
	machine.Memory[4] = 0x10
	machine.Memory[5] = 0x51
	machine.Memory[6] = 6
	machine.Memory[7] = 0x70

	machine.Run()

	fmt.Println(machine.RX)
}
