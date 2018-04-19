package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"bufio"
	"os"
)

var (
	// рабочие регистры
	x float64
	y float64

	// регистры общего назначеия
	r = map[string]float64{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
		"7": 0,
		"8": 0,
		"9": 0,
		"A": 0,
		"B": 0,
		"C": 0,
		"D": 0,
	}

	debug = true
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			continue
		}

		input = strings.Trim(input, "\n")

		if input == "" {
			continue
		}

		switch input {
		case "debug":
			debug = !debug
		case "dump":
			dump()
		case "cx":
			println("clear")
			x = 0
		case "<>":
			println("swap register x(", x, ") and register y(", y, ")")
			x, y = y, x
		case "-", "+", "/", "*":
			arithmetic(input)
		default:
			switch true {
			case strings.HasPrefix(input, "p"):
				parts := strings.Split(input, " ")
				name := parts[1]

				println("move x(", x, ") to r", name, "(", r[name], ")")
				r[name] = x
			case strings.HasPrefix(input, "ip"):
				parts := strings.Split(input, " ")
				name := parts[1]

				println("move r", name, "(", r[name], ") to x(", x, ")")
				x = r[name]
			default:
				y = x
				x, err = strconv.ParseFloat(input, 64)
				if err != nil {
					log.Println(err)
					continue
				}

				println("load", x, "to register x")
			}
		}

		fmt.Println("> ", x)
	}
}

func arithmetic(input string) {
	switch input {
	case "+":
		println("register y(", y, ") + register x(", x, ")")
		x = y + x
	case "-":
		println("register y(", y, ") - register x(", x, ")")
		x = y - x
	case "/":
		println("register y(", y, ") / register x(", x, ")")
		x = y / x
	case "*":
		println("register y(", y, ") * register x(", x, ")")
		x = y * x
	}
}

func println(args ...interface{}) {
	if debug {
		fmt.Println(args...)
	}
}

func dump() {
	fmt.Println("x =", x)
	fmt.Println("y =", y, "\n")
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
