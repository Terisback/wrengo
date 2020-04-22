package main

import (
	"fmt"

	"github.com/Terisback/wrengo"
)

func main() {
	program := `
		class WrenMath {
			static do_add(a, b) {
				return a + b
			}
			static do_sub(a, b) {
				return a - b
			}
			static do_mul(a, b) {
				return a * b
			}
			static do_div(a, b) {
				return a / b
			}
		}
	`

	// New configuration for VM
	config := wrengo.NewConfiguration()

	// Adding callbacks
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError

	// Creating new VM
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	// Interpret
	err := vm.Interpret(wrengo.DefaultModule, program)
	if err != nil {
		panic("Something went wrong")
	}

	// HANDLES!!!!!!!!! YAAAAAAAAY
	{
		vm.EnsureSlots(3)
		vm.GetVariable(wrengo.DefaultModule, "WrenMath", 0)
		h := vm.NewCallHandle("do_add(_,_)")
		vm.SetSlotDouble(1, 9)
		vm.SetSlotDouble(2, 3)
		err = h.Call()
		if err != nil {
			panic("Something went wrong")
		}
		fmt.Println(vm.GetSlotDouble(0))
	}

	{
		vm.EnsureSlots(3)
		vm.GetVariable(wrengo.DefaultModule, "WrenMath", 0)
		h := vm.NewCallHandle("do_sub(_,_)")
		vm.SetSlotDouble(1, 9)
		vm.SetSlotDouble(2, 3)
		err = h.Call()
		if err != nil {
			panic("Something went wrong")
		}
		fmt.Println(vm.GetSlotDouble(0))
	}

	{
		vm.EnsureSlots(3)
		vm.GetVariable(wrengo.DefaultModule, "WrenMath", 0)
		h := vm.NewCallHandle("do_mul(_,_)")
		vm.SetSlotDouble(1, 9)
		vm.SetSlotDouble(2, 3)
		err = h.Call()
		if err != nil {
			panic("Something went wrong")
		}
		fmt.Println(vm.GetSlotDouble(0))
	}

	{
		vm.EnsureSlots(3)
		vm.GetVariable(wrengo.DefaultModule, "WrenMath", 0)
		h := vm.NewCallHandle("do_div(_,_)")
		vm.SetSlotDouble(1, 9)
		vm.SetSlotDouble(2, 3)
		err = h.Call()
		if err != nil {
			panic("Something went wrong")
		}
		fmt.Println(vm.GetSlotDouble(0))
	}
}
