package executor

import (
	"errors"

	"github.com/horechek/bz34/vm"
	"github.com/horechek/bz34/vm/commands"
)

type Executor struct {
	vm *vm.VirtualMachine
}

func NewExecutor(vm *vm.VirtualMachine) *Executor {
	return &Executor{
		vm: vm,
	}
}

// выполняем команду, приводим виртуальную машину к определенному состоянию
func (e *Executor) Execute(code int) error {
	if code == 0 {
		return nil
	}

	cmd, ok := commands.Commands[code]
	if !ok {
		return errors.New("undefined code")
	}

	cmd.Execute(e.vm)

	return nil
}
