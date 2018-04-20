package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserStack(t *testing.T) {
	prepare()
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
		parser(input)
	}

	assert.Equal(t, 119.42742497399934, x)
}

func TestParserFunctions(t *testing.T) {
	prepare()
	inputs := []string{
		"754",
		":sqrt",
		"p 1",
		"4",
		"9.81",
		":x^y",
		"4.31",
		":x^2",
		"/",
		"ip 1",
		"+",
		"p 1",
		"6.31",
		"0.532",
		"*",
		"p 2",
		"5.81",
		"8.53",
		":sqrt",
		"*",
		"ip 2",
		"+",
		"ip 1",
		"xy",
		"/",
	}

	for _, input := range inputs {
		parser(input)
	}

	assert.Equal(t, 25.87973300826302, x)
}

// чистим все рагистры перед запуском
func prepare() {
	x = 0
	y = 0
	z = 0
	t = 0
	x1 = 0
	grad = false
	d = false

	for name := range r {
		r[name] = 0
	}
}
