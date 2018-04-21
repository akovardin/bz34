package vm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrigonometricSin(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"0.754",
		"fsin",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 0.6845600545913451, vm.Result())
}

func TestTrigonometricCos(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"fpi",
		"4",
		"*",
		"7",
		"/",
		"fcos",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, -0.22252093395631434, vm.Result())
}

func TestTrigonometricTg(t *testing.T) {
	vm := NewVirtualMachine()
	inputs := []string{
		"fpi",
		"2",
		"*",
		"7",
		"/",
		"ftg",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 1.2539603376627035, vm.Result())
}

func TestTrigonometricArctg(t *testing.T) {
	vm := NewVirtualMachine()
	vm.grad = true // ответ должен быть в градусах

	inputs := []string{
		"41",
		"farctg",
	}

	for _, input := range inputs {
		code, err := vm.Parse(input)
		assert.NoError(t, err)
		err = vm.Execute(code)
		assert.NoError(t, err)
	}

	assert.Equal(t, 88.60281897270363, vm.Result())
}
