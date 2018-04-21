package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/horechek/bz34/vm"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	machine := vm.NewVirtualMachine()
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

		code, err := machine.Parse(input)
		if err != nil {
			log.Println(err)
			continue
		}

		err = machine.Execute(code)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("> ", machine.Result())
	}
}
