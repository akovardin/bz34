package parser

import (
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
	// remove this - it is part of executor
	val, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		return 0, err
	}

	code, ok := tokens[input]
	if !ok {
		return int(val), nil
	}

	return code, nil
}
