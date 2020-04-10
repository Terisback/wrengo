package main

/*
#include "wren.h"
*/
import "C"
import (
	"fmt"
)

type InterpretResult int

const (
	RESULT_SUCCESS InterpretResult = iota
	RESULT_COMPILE_ERROR
	RESULT_RUNTIME_ERROR
)

func (i InterpretResult) String() string {
	switch i {
	case RESULT_SUCCESS:
		return "RESULT_SUCCESS"
	case RESULT_COMPILE_ERROR:
		return "RESULT_COMPILE_ERROR"
	case RESULT_RUNTIME_ERROR:
		return "RESULT_RUNTIME_ERROR"
	default:
		return ""
	}
}

type VM struct {
	wren *C.WrenVM
}

func NewVM(config *C.WrenConfiguration) VM {
	vm := VM{}
	vm.wren = C.wrenNewVM(config)
	return vm
}

func (vm VM) Interpret(module, source string) InterpretResult {
	m, s := C.CString(module), C.CString(source)
	defer C.free(unsafe.Pointer(m))
	defer C.free(unsafe.Pointer(s))
	return InterpretResult(C.wrenInterpret(vm.wren, m, s))
}

func main() {
	config := C.WrenConfiguration{}
	C.wrenInitConfiguration(&config)
	vm := NewVM(&config)

	fmt.Println(vm.Interpret("my_module", "System.print(\"I am running in a VM!\")").String())
}
