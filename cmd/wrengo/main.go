package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Terisback/wrengo"
)

func main() {
	var dir string

	if len(os.Args) == 2 {
		dir = os.Args[1]
	}

	// New configuration for VM
	config := wrengo.NewConfiguration()

	// Adding callbacks
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError

	// Creating new VM
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	if dir != "" {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println(dir, "not exist")
			os.Exit(1)
		}

		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			code, err := ioutil.ReadFile(dir)
			if err != nil {
				fmt.Println("idk why it's doesn't work, but the error is", err)
				os.Exit(1)
			}
			err = vm.Interpret(wrengo.DefaultModule, string(code))
			if err != nil {
				fmt.Println("Interpret says it don't like you, and also it says", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

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
