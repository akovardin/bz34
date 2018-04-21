package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser10powXFunctions(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"0.9374",
		"f10^x",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 8.657649506603732, vm.Result())
}

func TestParserEpowXFunctions(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"3.2705",
		"fe^x",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 26.324498302403168, vm.Result())
}

func TestParserLgFunctions(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"0.37445",
		"flgx",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, -0.4266061650774793, vm.Result())

}

func TestParserLnFunctions(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"0.005776",
		"flnx",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, -5.1540438773916115, vm.Result())
}
