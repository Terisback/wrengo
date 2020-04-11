package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Terisback/wrengo"
)

func main() {
	// Creating new VM
	config := wrengo.NewConfiguration()
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	fmt.Println("Ready! Print \"quit\" or Ctrl+C for exit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		if text == "quit" {
			break
		}

		vm.Interpret("interpret", text+"\n")
	}
}
