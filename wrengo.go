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

func main() {
	config := C.WrenConfiguration{}
	C.wrenInitConfiguration(&config)
	vm := C.wrenNewVM(&config)
	module := C.CString("my_module")
	code := C.CString("System.print(\"I am running in a VM!\")")
	result := C.wrenInterpret(vm, module, code)
	switch InterpretResult(result) {
	case RESULT_SUCCESS:
		fmt.Println("RESULT_SUCCESS")
	case RESULT_COMPILE_ERROR:
		fmt.Println("RESULT_COMPILE_ERROR")
	case RESULT_RUNTIME_ERROR:
		fmt.Println("RESULT_RUNTIME_ERROR")
	}
}
