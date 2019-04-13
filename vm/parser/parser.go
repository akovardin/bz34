package parser

import (
	"fmt"
	"strconv"

	"github.com/horechek/bz34/vm"
)

type Parser struct {
	vm *vm.VirtualMachine
}

func NewParser(vm *vm.VirtualMachine) *Parser {
	return &Parser{
		vm: vm,
	}
}

func (p *Parser) Parse(input string) (int, error) {
	val, err := strconv.ParseFloat(input, 64)
	if err == nil {
		p.vm.Save()
		p.vm.Put()
		p.vm.RX = val
		return 0, nil // enter code
	}

	if input == "dump" {
		p.vm.Dump()
		return 0, nil
	}

	code, ok := tokens[input]
	if !ok {
		return 0, fmt.Errorf("undefined token %s", input)
	}

	return code, nil
}
