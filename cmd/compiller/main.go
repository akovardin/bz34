// compile text file to binary code
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/horechek/bz34/vm/parser"
)

func main() {
	parser := parser.NewParser()
	prog := "1 2 +"

	data := strings.Split(prog, " ")
	codes := []int{}
	for _, c := range data {
		code, err := parser.Parse(c)
		if err != nil {
			log.Fatal(err)
		}

		codes = append(codes, code)
	}

	fmt.Println(codes) // write codes to binary file
}
