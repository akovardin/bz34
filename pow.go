package main

import "math"

func pow(input string) {
	switch input {
	case ":sqrt":
		x1 = x
		x = math.Sqrt(x)
	case ":x^2":
		x1 = x
		x = math.Pow(x, 2)
	case ":x^y":
		x1 = x
		x = math.Pow(x, y)
	case ":1/x":
		x1 = x
		x = 1 / x
	}
}
