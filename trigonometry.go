package main

import (
	"math"
)

// переключатель радианы/градусы
var (
	grad = false
)

func trigonometry(input string) {
	x1 = x
	switch input {
	case ":pi":
		x = math.Pi
	case ":sin":
		x = math.Sin(x)
	case ":cos":
		x = math.Cos(x)
	case ":tg":
		x = math.Tan(x)
	case ":ctg":
		x = math.Pow(math.Tanh(x), -1)
	case ":arcsin":
		x = math.Asin(x)
	case ":arccos":
		x = math.Acos(x)
	case ":arctg":
		x = math.Atan(x)
	case ":arcctg":
		x = math.Pow(math.Atan(x), -1)
	}

	if grad {
		x = x / (math.Pi / 180)
	}
}
