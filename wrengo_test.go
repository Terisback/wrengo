package wrengo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileError(t *testing.T) {
	vm := NewVM(NewConfiguration())
	defer vm.FreeVM()

	err := vm.Interpret(DefaultModule, `Hay, it's me, an error!`)
	if err == nil {
		t.FailNow()
	}
}

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

func TestSlots(t *testing.T) {
	vm := NewVM(NewConfiguration())
	defer vm.FreeVM()

	vm.EnsureSlots(1)
	assert.Equal(t, 1, vm.GetSlotCount())

	vm.SetSlotBool(0, true)
	assert.Equal(t, true, vm.GetSlotBool(0))

	vm.SetSlotBytes(0, []byte("Hello"))
	assert.Equal(t, []byte("Hello"), vm.GetSlotBytes(0, len([]byte("Hello"))))

	vm.SetSlotDouble(0, 1.337)
	assert.Equal(t, 1.337, vm.GetSlotDouble(0))

	vm.SetSlotString(0, "Hi")
	assert.Equal(t, "Hi", vm.GetSlotString(0))
}

type God struct {
	msg string
}

// The constructor for a foreign type takes no arguments and returns
// an interface{} value representing the new object.
func NewGod() interface{} {
	return &God{msg: "What are you doing? "}
}

func GetGodsMessage(vm *VM) {
	god := vm.GetSlotForeign(0, God{}).(God)
	name := vm.GetSlotString(1)
	vm.SetSlotString(0, god.msg+name)
}

func TestForeign(t *testing.T) {
	program := `
			foreign class God {
				construct new() {}
				foreign getMessage(name)
			}

			var god = God.new()
			System.print(god.getMessage("Damien"))
		`

	var out string

	config := NewConfiguration()
	config.WriteFunc = func(vm *VM, text string) {
		out += text
	}
	config.ErrorFunc = CallbackError
	vm := NewVM(config)
	assert.NoError(t, vm.BindForeignClass("God", NewGod))
	assert.NoError(t, vm.BindForeignMethod("God", false, "getMessage(_)", GetGodsMessage))

	assert.NoError(t, vm.Interpret(DefaultModule, program))

	assert.Equal(t, "What are you doing? Damien\n", out)
}
