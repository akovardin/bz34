package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser10powXFunctions(t *testing.T) {
	prepare()
	inputs := []string{
		"0.9374",
		":10^x",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 8.657649506603732, x)
}

func TestParserEpowXFunctions(t *testing.T) {
	prepare()
	inputs := []string{
		"3.2705",
		":e^x",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 26.324498302403168, x)
}

func TestParserLgFunctions(t *testing.T) {
	prepare()
	inputs := []string{
		"0.37445",
		":lgx",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, -0.4266061650774793, x)

}

func TestParserLnFunctions(t *testing.T) {
	prepare()
	inputs := []string{
		"0.005776",
		":lnx",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, -5.1540438773916115, x)
}
