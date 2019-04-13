package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/horechek/bz34/vm"
	"github.com/horechek/bz34/vm/executor"
	"github.com/horechek/bz34/vm/parser"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	machine := vm.NewVirtualMachine()
	reader := bufio.NewReader(os.Stdin)
	parser := parser.NewParser(machine)
	executor := executor.NewExecutor(machine)

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

		code, err := parser.Parse(input)
		if err != nil {
			log.Println(err)
			continue
		}

		err = executor.Execute(code)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("> ", machine.Result())
	}
}
