package main

func arithmetic(input string) {
	x1 = x

	switch input {
	case "+":
		println("register y(", y, ") + register x(", x, ")")
		y = y + x
	case "-":
		println("register y(", y, ") - register x(", x, ")")
		y = y - x
	case "/":
		println("register y(", y, ") / register x(", x, ")")
		y = y / x
	case "*":
		println("register y(", y, ") * register x(", x, ")")
		y = y * x
	}

	pop()
}
