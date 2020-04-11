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

	fmt.Println("Ready! Press Ctrl+C for exit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		vm.Interpret("main", text+"\n")
	}
}
