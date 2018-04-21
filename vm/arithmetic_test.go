package vm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleArithmeticStack(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"35",
		"7",
		"-",
		"26",
		"12",
		"-",
		"/",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.EqualValues(t, 2, vm.Result())
}

func TestArithmeticStack(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"27.41",
		"13.91",
		"-",
		"18.49",
		"17.403",
		"+",
		"*",
		"596.3",
		"585.49",
		"-",
		"*",
		"0.2597",
		"0.2369",
		"-",
		"*",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 119.42742497399934, vm.Result())
}
