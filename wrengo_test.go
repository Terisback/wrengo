package wrengo

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	err := vm.Interpret(DefaultModule, `System.print("Hello there!")`)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, "Hello there!\n", out)
}

func TestCallHandle(t *testing.T) {
	config := NewConfiguration()
	config.WriteFunc = CallbackWrite
	config.ErrorFunc = CallbackError
	vm := NewVM(config)
	defer vm.FreeVM()

	err := vm.Interpret(DefaultModule, `
		class WrenMath {
			static do_add(a, b) {
				return a + b
			}
		}
	`)
	if err != nil {
		t.FailNow()
	}

	vm.EnsureSlots(3)
	vm.GetVariable(DefaultModule, "WrenMath", 0)
	h := vm.NewCallHandle("do_add(_,_)")
	vm.SetSlotDouble(1, 228)
	vm.SetSlotDouble(2, 1337)
	err = h.Call()
	if err != nil {
		t.FailNow()
	}

	res := vm.GetSlotDouble(0)

	assert.Equal(t, float64(228+1337), res)
}
