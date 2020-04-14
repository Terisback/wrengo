package wrengo

import (
	"testing"
)

func TestInterpret(t *testing.T) {
	var out string

	config := NewConfiguration()
	config.WriteFunc = func(vm *VM, text string) {
		out += text
	}
	config.ErrorFunc = CallbackError
	vm := NewVM(config)
	defer vm.FreeVM()

	result := vm.Interpret(DefaultModule, `System.print("Hello there!")`)

	if out != "Hello there!\n" {
		t.Errorf(`Interpret("Hello there!") returned unexpected value: %s`, out)
	}

	if result != RESULT_SUCCESS {
		t.Errorf(`Interpret("Hello there!") have done with unexpected result: %s`, result.String())
	}
}

func TestVariables(t *testing.T) {
	config := NewConfiguration()
	config.WriteFunc = CallbackWrite
	config.ErrorFunc = CallbackError
	vm := NewVM(config)
	defer vm.FreeVM()

	result := vm.Interpret(DefaultModule, `
		class WrenMath {
			static do_add(a, b) {
				return a + b
			}
		}
	`)

	if result != RESULT_SUCCESS {
		t.Errorf(`TestVariables have done with unexpected result: %s`, result.String())
	}

	vm.EnsureSlots(3)
	vm.GetVariable(DefaultModule, "WrenMath", 0)
	h := vm.NewCallHandle("do_add(_,_)")
	vm.SetSlotDouble(1, 228)
	vm.SetSlotDouble(2, 1337)
	result = h.Call()

	if result != RESULT_SUCCESS {
		t.Errorf(`TestVariables method have done with unexpected result: %s`, result.String())
	}

	res := vm.GetSlotDouble(0)

	if res != 228+1337 {
		t.Errorf(`TestVariables method returned unexpected value: %s`, result.String())
	}
}
