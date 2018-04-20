package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	x float64 // текущий регистр
	y float64 // рабочий регистр
	// стековые регистры
	z float64
	t float64

	// регистр предыдущего результата
	x1 float64

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

	d = true
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

		parser(input)

		fmt.Println("> ", x)
	}
}

func println(args ...interface{}) {
	if d {
		fmt.Println(args...)
	}
}
