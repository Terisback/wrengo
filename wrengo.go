package wrengo

/*
#include "wren.h"

extern void wrengoWrite(WrenVM*, char*);
extern void wrengoError(WrenVM*, WrenErrorType, char*, int, char* );
*/
import "C"
import (
	"fmt"
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

	// The callback Wren uses to display text when `System.print()`
	// or the other related functions are called.
	//
	// If this is `NULL`, Wrengo by default writes into Stdout.
	WriteFunc func(vm *VM, text string)

	// The callback Wren uses to report errors.
	//
	// When an error occurs, this will be called with the module name, line
	// number, and an error message. If this is `NULL`,
	// Wrengo by default writes errors into Stdout.
	ErrorFunc func(vm *VM, errorType ErrorType, module string, line int, message string)
}

type Configuration struct {
	Callbacks

	// The number of bytes Wren will allocate before triggering the first garbage
	// collection.
	//
	// If zero, defaults to 10MB.
	InitialHeapSize uint

	// After a collection occurs, the threshold for the next collection is
	// determined based on the number of bytes remaining in use. This allows Wren
	// to shrink its memory usage automatically after reclaiming a large amount
	// of memory.
	//
	// This can be used to ensure that the heap does not get too small, which can
	// in turn lead to a large number of collections afterwards as the heap grows
	// back to a usable size.
	//
	// If zero, defaults to 1MB.
	MinHeapSize uint

	// Wren will resize the heap automatically as the number of bytes
	// remaining in use after a collection changes. This number determines the
	// amount of additional memory Wren will use after a collection, as a
	// percentage of the current heap size.
	//
	// For example, say that this is 50. After a garbage collection, when there
	// are 400 bytes of memory still in use, the next collection will be triggered
	// after a total of 600 bytes are allocated (including the 400 already in
	// use.)
	//
	// Setting this to a smaller number wastes less memory, but triggers more
	// frequent garbage collections.
	//
	// If zero, defaults to 50.
	HeapGrowthPercent int

	config *C.WrenConfiguration
}

func NewConfiguration() Configuration {
	cfg := Configuration{}
	cfg.config = &C.WrenConfiguration{}
	C.wrenInitConfiguration(cfg.config)
	cfg.WriteFunc = defaultWrite
	cfg.ErrorFunc = defaultError
	return cfg
}

// A single virtual machine for executing Wren code.
//
// Wren has no global state, so all state stored by a running interpreter lives
// here.
type VM struct {
	cb Callbacks

	vm *C.WrenVM
}

// Creates a new Wren virtual machine using the given [configuration].
func NewVM(cfg Configuration) VM {
	cfg.config.initialHeapSize = C.size_t(cfg.InitialHeapSize)
	cfg.config.minHeapSize = C.size_t(cfg.MinHeapSize)
	cfg.config.heapGrowthPercent = C.int(cfg.HeapGrowthPercent)

	cfg.config.writeFn = C.WrenWriteFn(C.wrengoWrite)
	cfg.config.errorFn = C.WrenErrorFn(C.wrengoError)

	vm := VM{}
	vm.vm = C.wrenNewVM(cfg.config)
	vm.cb = cfg.Callbacks
	vmMap[vm.vm] = &vm
	return vm
}

// Disposes of all resources is use by [vm], which was previously created by a
// call to [NewVM].
func (vm *VM) FreeVM() {
	C.wrenFreeVM(vm.vm)
	delete(vmMap, vm.vm)
}

// Immediately run the garbage collector to free unused memory.
func (vm *VM) GC() {
	C.wrenCollectGarbage(vm.vm)
}

// Runs [source], a string of Wren source code in a new fiber in VM in the
// context of resolved [module].
func (vm *VM) Interpret(module, source string) InterpretResult {
	return InterpretResult(C.wrenInterpret((*vm).vm, C.CString(module), C.CString(source)))
}

func defaultWrite(vm *VM, text string) {
	fmt.Print(text)
}

func defaultError(vm *VM, errorType ErrorType, module string, line int, message string) {
	fmt.Println(fmt.Sprint(errorType.String(), " (", module, ") Line ", line, " : ", message))
}

//export wrengoWrite
func wrengoWrite(vm *C.WrenVM, text *C.char) {
	vmMap[vm].cb.WriteFunc(vmMap[vm], C.GoString(text))
}

//export wrengoError
func wrengoError(vm *C.WrenVM, err C.WrenErrorType, module *C.char, line C.int, message *C.char) {
	vmMap[vm].cb.ErrorFunc(vmMap[vm], ErrorType(err), C.GoString(module), int(line), C.GoString(message))
}
