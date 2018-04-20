package main

func stack(input string) {
	switch input {
	case ":pop":
		pop()
	case ":put":
		put()
	case ":bx":
		bx()
	}
}

func bx() {
	y = x
	x = x1
}

func put() {
	t = z
	z = y
	y = x
}

func pop() {
	x = y
	y = z
	z = t
}
