package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrigonometricSin(t *testing.T) {
	prepare()
	inputs := []string{
		"0.754",
		":sin",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 0.6845600545913451, x)
}

func TestTrigonometricCos(t *testing.T) {
	prepare()
	inputs := []string{
		":pi",
		"4",
		"*",
		"7",
		"/",
		":cos",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, -0.22252093395631434, x)
}

func TestTrigonometricTg(t *testing.T) {
	prepare()
	inputs := []string{
		":pi",
		"2",
		"*",
		"7",
		"/",
		":tg",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 1.2539603376627035, x)
}

func TestTrigonometricArctg(t *testing.T) {
	prepare()
	grad = true // ответ должен быть в градусах

	inputs := []string{
		"41",
		":arctg",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 88.60281897270363, x)
}
