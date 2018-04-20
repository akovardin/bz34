package main

import "math"

func ten(input string) {
	x1 = x

	switch input {
	case ":e^x":
		x = math.Pow(math.E, x)
	case ":lgx":
		x = math.Log10(x)
	case ":lnx":
		x = math.Log(x)
	case ":10^x":
		x = math.Pow(10, x)
	}
}
