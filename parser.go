package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parser(input string) error {
	switch true {
	case strings.HasPrefix(input, "debug"):
		debug()
	case strings.HasPrefix(input, "dump"):
		dump()
	case strings.HasPrefix(input, "cx"):
		clear()
	case strings.HasPrefix(input, "xy"):
		swap()
	case strings.HasPrefix(input, "-"),
		strings.HasPrefix(input, "+"),
		strings.HasPrefix(input, "/"),
		strings.HasPrefix(input, "*"):
		arithmetic(input)
	case strings.HasPrefix(input, "p"):
		save(input)
	case strings.HasPrefix(input, "ip"):
		load(input)
	case strings.HasPrefix(input, "fn"):
		function(input)
	default:
		put()
		var err error
		x, err = strconv.ParseFloat(input, 64)
		if err != nil {
			return err
		}
		println("load", x, "to register x")
	}

	return nil
}

func debug() {
	d = !d
}

func dump() {
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
	fmt.Println("t =", t)
	fmt.Println()

	fmt.Println("x1 =", x1)
	fmt.Println()

	fmt.Println("0 =", r["0"])
	fmt.Println("1 =", r["1"])
	fmt.Println("2 =", r["2"])
	fmt.Println("3 =", r["3"])
	fmt.Println("4 =", r["4"])
	fmt.Println("5 =", r["5"])
	fmt.Println("6 =", r["6"])
	fmt.Println("7 =", r["7"])
	fmt.Println("8 =", r["8"])
	fmt.Println("9 =", r["9"])
	fmt.Println("A =", r["A"])
	fmt.Println("B =", r["B"])
	fmt.Println("C =", r["C"])
	fmt.Println("D =", r["D"])
}

func clear() {
	println("clear")
	x = 0
}

func swap() {
	println("swap register x(", x, ") and register y(", y, ")")
	x, y = y, x
}

func save(input string) {
	parts := strings.Split(input, " ")
	name := parts[1]
	println("move x(", x, ") to r", name, "(", r[name], ")")
	r[name] = x
}

func load(input string) {
	parts := strings.Split(input, " ")
	name := parts[1]
	println("move r", name, "(", r[name], ") to x(", x, ")")

	put() // передвигаем по стеку

	x = r[name]
}

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

func function(input string) {
	parts := strings.Split(input, " ")
	name := parts[1]
	println("call  function", name)

	switch name {
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
	case ":bx":
		y = x
		x = x1
	case ":pop":
		pop()
	}
}

// move to stack functions
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
