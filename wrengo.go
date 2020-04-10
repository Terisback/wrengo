package main

/*
#include "wren.h"

extern void wrengoWrite(WrenVM*, char*);
extern void wrengoError(WrenVM*, WrenErrorType, char*, int, char* );
*/
import "C"
import (
	"fmt"
	"runtime"
)

var (
	vmMap = make(map[*C.WrenVM]*VM)
)

type InterpretResult int

const (
	RESULT_SUCCESS InterpretResult = iota
	RESULT_COMPILE_ERROR
	RESULT_RUNTIME_ERROR
)

func (i InterpretResult) String() string {
	return [...]string{"RESULT_SUCCESS", "RESULT_COMPILE_ERROR", "RESULT_RUNTIME_ERROR"}[i]
}

type ErrorType int

const (
	// A syntax or resolution error detected at compile time.
	ERROR_COMPILE ErrorType = iota

	// The error message for a runtime error.
	ERROR_RUNTIME

	// One entry of a runtime error's stack trace.
	ERROR_STACK_TRACE
)

func (i ErrorType) String() string {
	return [...]string{"ERROR_COMPILE", "ERROR_RUNTIME", "ERROR_STACK_TRACE"}[i]
}

type Callbacks struct {
	WriteFunc func(vm *VM, text string)
	ErrorFunc func(vm *VM, errorType ErrorType, module string, line int, message string)
}

type Configuration struct {
	Callbacks

	config *C.WrenConfiguration
}

func NewConfiguration() Configuration {
	cfg := Configuration{}
	cfg.config = &C.WrenConfiguration{}
	C.wrenInitConfiguration(cfg.config)
	cfg.WriteFunc = func(vm *VM, text string) {
		fmt.Print(text)
	}
	cfg.ErrorFunc = func(vm *VM, errorType ErrorType, module string, line int, message string) {
		fmt.Println("Error happend:", errorType.String(), "in", module, "module\n", line, ":", message)
	}
	return cfg
}

type VM struct {
	cb Callbacks

	vm *C.WrenVM
}

func NewVM(cfg Configuration) VM {
	cfg.config.writeFn = C.WrenWriteFn(C.wrengoWrite)
	cfg.config.errorFn = C.WrenErrorFn(C.wrengoError)

	vm := VM{}
	vm.vm = C.wrenNewVM(cfg.config)
	vm.cb = cfg.Callbacks
	vmMap[vm.vm] = &vm
	runtime.SetFinalizer(&vm, func(vm *VM) {
		C.wrenFreeVM(vm.vm)
		delete(vmMap, vm.vm)
	})
	return vm
}

func (vm *VM) Interpret(module, source string) InterpretResult {
	return InterpretResult(C.wrenInterpret((*vm).vm, C.CString(module), C.CString(source)))
}

func main() {
	config := NewConfiguration()
	vm := NewVM(config)
	fmt.Println(vm.Interpret("my_module", "System.print(\"I am running in a VM!\")").String())
}

//export wrengoWrite
func wrengoWrite(vm *C.WrenVM, text *C.char) {
	vmMap[vm].cb.WriteFunc(vmMap[vm], C.GoString(text))
}

//export wrengoError
func wrengoError(vm *C.WrenVM, err C.WrenErrorType, module *C.char, line C.int, message *C.char) {
	vmMap[vm].cb.ErrorFunc(vmMap[vm], ErrorType(err), C.GoString(module), int(line), C.GoString(message))
}
