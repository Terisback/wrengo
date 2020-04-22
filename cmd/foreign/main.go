package main

import (
	"fmt"
	"log"

	"github.com/Terisback/wrengo"
)

type God struct {
	msg string
}

// The constructor for a foreign type takes no arguments and returns
// an interface{} value representing the new object.
func NewGod() interface{} {
	return &God{msg: "What are you doing? %s"}
}

func GetGodsMessage(vm *wrengo.VM) {
	god := vm.GetSlotForeign(0, God{}).(God)
	name := vm.GetSlotString(1)
	vm.SetSlotString(0, fmt.Sprintf(god.msg, name))
}

func main() {
	// Wren code
	program := `
			foreign class God {
				construct new() {}
				foreign getMessage(name)
			}

			var god = God.new()
			System.print(god.getMessage("Silly boy"))
		`

	// New configuration for VM
	config := wrengo.NewConfiguration()

	// Adding callbacks
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError

	// Creating new VM
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	// Bind God class
	err := vm.BindForeignClass("God", NewGod)
	if err != nil {
		log.Fatalln(err)
	}

	// Bind God.getMessage()
	err := vm.BindForeignMethod("God", false, "getMessage(_)", GetGodsMessage)
	if err != nil {
		log.Fatalln(err)
	}

	// Interpret
	err := vm.Interpret(wrengo.DefaultModule, program)
	if err != nil {
		log.Fatalln(err)
	}
}