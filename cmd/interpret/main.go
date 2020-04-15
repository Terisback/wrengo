package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Terisback/wrengo"
)

func main() {
	// New configuration for VM
	config := wrengo.NewConfiguration()

	// Adding callbacks
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError

	// Creating new VM
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	// New console read buffer
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ready! Press Ctrl+C for exit.")

	// Interpret loop
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		err := vm.Interpret("main", text+"\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}
