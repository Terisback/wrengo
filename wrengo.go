package wrengo

/*
#cgo LDFLAGS: -L${SRCDIR}/lib -lwren
#include "wren.h"

extern char* wrengoResolveModule(WrenVM*, char*, char*);
extern char* wrengoLoadModule(WrenVM*, char*);
extern void* wrengoBindForeignMethod(WrenVM*, char*, char*, bool, char*);
extern WrenForeignClassMethods wrengoBindForeignClass(WrenVM*, char*, char*);
extern void wrengoWrite(WrenVM*, char*);
extern void wrengoError(WrenVM*, WrenErrorType, char*, int, char* );
*/
import "C"
import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

const (
	DefaultModule string = "main"
)

var (
	vmMap = make(map[*C.WrenVM]*VM)
)

type Callbacks struct {

	// The callback Wren uses to resolve a module name.
	//
	// Some host applications may wish to support "relative" imports, where the
	// meaning of an import string depends on the module that contains it. To
	// support that without baking any policy into Wren itself, the VM gives the
	// host a chance to resolve an import string.
	//
	// Before an import is loaded, it calls this, passing in the name of the
	// module that contains the import and the import string. The host app can
	// look at both of those and produce a new "canonical" string that uniquely
	// identifies the module. This string is then used as the name of the module
	// going forward. It is what is passed to [loadModuleFn], how duplicate
	// imports of the same module are detected, and how the module is reported in
	// stack traces.
	//
	// If you leave this function NULL, then the original import string is
	// treated as the resolved string.
	//
	// If an import cannot be resolved by the embedder, it should return NULL and
	// Wren will report that as a runtime error.
	//
	// Wren will take ownership of the string you return and free it for you, so
	// it should be allocated using the same allocation function you provide
	// above.
	ResolveModuleFunc func(vm *VM, importer, name string) string

	// The callback Wren uses to load a module.
	//
	// Since Wren does not talk directly to the file system, it relies on the
	// embedder to physically locate and read the source code for a module. The
	// first time an import appears, Wren will call this and pass in the name of
	// the module being imported. The VM should return the soure code for that
	// module. Memory for the source should be allocated using [reallocateFn] and
	// Wren will take ownership over it.
	//
	// This will only be called once for any given module name. Wren caches the
	// result internally so subsequent imports of the same module will use the
	// previous source and not call this.
	//
	// If a module with the given name could not be found by the embedder, it
	// should return NULL and Wren will report that as a runtime error.
	LoadModuleFunc func(vm *VM, name string) string

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
	return cfg
}

// A single virtual machine for executing Wren code.
//
// Wren has no global state, so all state stored by a running interpreter lives
// here.
type VM struct {
	cb               Callbacks
	classes, methods map[string]unsafe.Pointer
	vm               *C.WrenVM
}

// Creates a new Wren virtual machine using the given [configuration].
func NewVM(cfg Configuration) VM {
	cfg.config.initialHeapSize = C.size_t(cfg.InitialHeapSize)
	cfg.config.minHeapSize = C.size_t(cfg.MinHeapSize)
	cfg.config.heapGrowthPercent = C.int(cfg.HeapGrowthPercent)

	cfg.config.bindForeignMethodFn = C.WrenBindForeignMethodFn(C.wrengoBindForeignMethod)
	cfg.config.bindForeignClassFn = C.WrenBindForeignClassFn(C.wrengoBindForeignClass)

	if cfg.ResolveModuleFunc != nil {
		cfg.config.resolveModuleFn = C.WrenResolveModuleFn(C.wrengoResolveModule)
	}

	if cfg.LoadModuleFunc != nil {
		cfg.config.loadModuleFn = C.WrenLoadModuleFn(C.wrengoLoadModule)
	}

	if cfg.WriteFunc != nil {
		cfg.config.writeFn = C.WrenWriteFn(C.wrengoWrite)
	}

	if cfg.ErrorFunc != nil {
		cfg.config.errorFn = C.WrenErrorFn(C.wrengoError)
	}

	vm := VM{}
	vm.vm = C.wrenNewVM(cfg.config)
	vm.classes = make(map[string]unsafe.Pointer)
	vm.methods = make(map[string]unsafe.Pointer)
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
func (vm *VM) Interpret(module, source string) error {
	m, s := C.CString(module), C.CString(source)
	defer C.free(unsafe.Pointer(m))
	defer C.free(unsafe.Pointer(s))
	return InterpretResult(C.wrenInterpret(vm.vm, m, s)).Error()
}

// A handle to a Wren object.
//
// This lets code outside of the VM hold a persistent reference to an object.
// After a handle is acquired, and until it is released, this ensures the
// garbage collector will not reclaim the object it references.
type Handle struct {
	vm     *VM
	handle *C.WrenHandle
}

// Creates a handle that can be used to invoke a method with [signature] on
// using a receiver and arguments that are set up on the stack.
//
// This handle can be used repeatedly to directly invoke that method from C
// code using [Call].
//
// When you are done with this handle, it must be released using
// [ReleaseHandle].
func (vm *VM) NewCallHandle(signature string) Handle {
	s := C.CString(signature)
	defer C.free(unsafe.Pointer(s))
	return Handle{vm: vm, handle: C.wrenMakeCallHandle(vm.vm, s)}
}

// Calls method, using the receiver and arguments previously set up on the
// stack.
//
// Method must have been created by a call to [NewCallHandle]. The
// arguments to the method must be already on the stack. The receiver should be
// in slot 0 with the remaining arguments following it, in order. It is an
// error if the number of arguments provided does not match the method's
// signature.
//
// After this returns, you can access the return value from slot 0 on the stack.
func (h *Handle) Call() error {
	return InterpretResult(C.wrenCall(h.vm.vm, h.handle)).Error()
}

// Releases the reference stored in [handle]. After calling this, [handle] can
// no longer be used.
func (h *Handle) Release() {
	C.wrenReleaseHandle(h.vm.vm, h.handle)
}

// Returns the number of slots available to the current foreign method.
func (vm *VM) GetSlotCount() int {
	return int(C.wrenGetSlotCount(vm.vm))
}

// Ensures that the foreign method stack has at least [numSlots] available for
// use, growing the stack if needed.
//
// Does not shrink the stack if it has more than enough slots.
//
// It is an error to call this from a finalizer.
func (vm *VM) EnsureSlots(numSlots int) {
	C.wrenEnsureSlots(vm.vm, C.int(numSlots))
}

// Gets the type of the object in [slot].
func (vm *VM) GetSlotType(slot int) WrenType {
	return WrenType(C.wrenGetSlotType(vm.vm, C.int(slot)))
}

// Reads a boolean value from [slot].
//
// It is an error to call this if the slot does not contain a boolean value.
func (vm *VM) GetSlotBool(slot int) bool {
	return bool(C.wrenGetSlotBool(vm.vm, C.int(slot)))
}

// Reads a byte array from [slot].
//
// The memory for the returned string is owned by Wren. You can inspect it
// while in your foreign method, but cannot keep a pointer to it after the
// function returns, since the garbage collector may reclaim it.
//
// Returns a pointer to the first byte of the array and fill [length] with the
// number of bytes in the array.
//
// It is an error to call this if the slot does not contain a string.
func (vm *VM) GetSlotBytes(slot, length int) []byte {
	l := C.int(length)
	data := C.GoString(C.wrenGetSlotBytes(vm.vm, C.int(slot), &l))
	return []byte(data)
}

// Reads a number from [slot].
//
// It is an error to call this if the slot does not contain a number.
func (vm *VM) GetSlotDouble(slot int) float64 {
	return float64(C.wrenGetSlotDouble(vm.vm, C.int(slot)))
}

// Reads a foreign object from [slot] and returns a pointer to the foreign data
// stored with it.
//
// It is an error to call this if the slot does not contain an instance of a
// foreign class.
func (vm *VM) GetSlotForeign(slot int, i interface{}) {
	ptr := C.wrenGetSlotForeign(vm.vm, C.int(slot)) // ptr
	i = reflect.NewAt(reflect.TypeOf(i).Elem(), ptr).Interface()
	// TODO: Implement return struct from ptr to i
	// i = ptr
}

// Reads a string from [slot].
//
// The memory for the returned string is owned by Wren. You can inspect it
// while in your foreign method, but cannot keep a pointer to it after the
// function returns, since the garbage collector may reclaim it.
//
// It is an error to call this if the slot does not contain a string.
func (vm *VM) GetSlotString(slot int) string {
	return C.GoString(C.wrenGetSlotString(vm.vm, C.int(slot)))
}

// Creates a handle for the value stored in [slot].
//
// This will prevent the object that is referred to from being garbage collected
// until the handle is released by calling [wrenReleaseHandle()].
func (vm *VM) GetSlotHandle(slot int) Handle {
	return Handle{vm: vm, handle: C.wrenGetSlotHandle(vm.vm, C.int(slot))}
}

// Stores the boolean [value] in [slot].
func (vm *VM) SetSlotBool(slot int, value bool) {
	C.wrenSetSlotBool(vm.vm, C.int(slot), C.bool(value))
}

// Stores the array [length] of [bytes] in [slot].
//
// The bytes are copied to a new string within Wren's heap, so you can free
// memory used by them after this is called.
func (vm *VM) SetSlotBytes(slot int, value []byte) {
	val := C.CString(string(value))
	defer C.free(unsafe.Pointer(val))
	C.wrenSetSlotBytes(vm.vm, C.int(slot), val, C.size_t(len(value)))
}

// Stores the numeric [value] in [slot].
func (vm *VM) SetSlotDouble(slot int, value float64) {
	C.wrenSetSlotDouble(vm.vm, C.int(slot), C.double(value))
}

// Stores a new empty list in [slot].
func (vm *VM) SetSlotNewList(slot int) {
	C.wrenSetSlotNewList(vm.vm, C.int(slot))
}

// Stores null in [slot].
func (vm *VM) SetSlotNull(slot int) {
	C.wrenSetSlotNull(vm.vm, C.int(slot))
}

// Stores the string [text] in [slot].
//
// The [text] is copied to a new string within Wren's heap, so you can free
// memory used by it after this is called. The length is calculated using
// [strlen()]. If the string may contain any null bytes in the middle, then you
// should use [wrenSetSlotBytes()] instead.
func (vm *VM) SetSlotString(slot int, value string) {
	val := C.CString(value)
	defer C.free(unsafe.Pointer(val))
	C.wrenSetSlotString(vm.vm, C.int(slot), val)
}

// Stores the value captured in [handle] in [slot].
//
// This does not release the handle for the value.
func (vm *VM) SetSlotHandle(slot int, handle Handle) {
	C.wrenSetSlotHandle(vm.vm, C.int(slot), handle.handle)
}

// Returns the number of elements in the list stored in [slot].
func (vm *VM) GetListCount(slot int) int {
	return int(C.wrenGetListCount(vm.vm, C.int(slot)))
}

// Reads element [index] from the list in [listSlot] and stores it in
// [elementSlot].
func (vm *VM) GetListElement(listSlot, index, elementSlot int) {
	C.wrenGetListElement(vm.vm, C.int(listSlot), C.int(index), C.int(elementSlot))
}

// Takes the value stored at [elementSlot] and inserts it into the list stored
// at [listSlot] at [index].
//
// As in Wren, negative indexes can be used to insert from the end. To append
// an element, use `-1` for the index.
func (vm *VM) InsertInList(listSlot, index, elementSlot int) {
	C.wrenInsertInList(vm.vm, C.int(listSlot), C.int(index), C.int(elementSlot))
}

// Looks up the top level variable with [name] in resolved [module] and stores
// it in [slot].
func (vm *VM) GetVariable(module, name string, slot int) {
	m, n := C.CString(module), C.CString(name)
	defer C.free(unsafe.Pointer(m))
	defer C.free(unsafe.Pointer(n))
	C.wrenGetVariable(vm.vm, m, n, C.int(slot))
}

// Sets the current fiber to be aborted, and uses the value in [slot] as the
// runtime error object.
func (vm *VM) AbortFiber(slot int) {
	C.wrenAbortFiber(vm.vm, C.int(slot))
}

// Registers a foreign method with the virtual machine.
func (vm *VM) BindForeignMethod(class string, isStatic bool, signature string, f func(*VM)) error {
	ptr, err := registerFunc(signature, f)
	if err != nil {
		return err
	}
	vmMap[vm.vm].methods[bindSignature(DefaultModule, class, isStatic, signature)] = ptr
	return nil
}

// Registers a foreign class with the virtual machine.
func (vm *VM) BindForeignClass(signature string, f func() interface{}) error {
	ptr, err := registerClass(signature, func() {
		newForeign(vm.vm, f())
	})
	if err != nil {
		return err
	}
	vmMap[vm.vm].classes[signature] = ptr
	return nil
}

func bindSignature(module, class string, isStatic bool, signature string) string {
	var sig bytes.Buffer
	if isStatic {
		sig.WriteString("static ")
	}
	sig.WriteString(class)
	sig.WriteString(".")
	sig.WriteString(signature)

	return sig.String()
}

// DON'T USE THIS FUNCTION - IT'S COPYPASTA FROM GO-WREN
// AFTER CALLING THUS FUNCTION CAN BE UNPREDICTABLE CONSEQUENCES
// newForeign allocates a new foreign object.
//
// This method should only be called from a foreign class allocation function.
// It takes an instance of the VM and a newly allocated foreign object ("foreign"
// meaning that it's created in Go and not Wren) and makes it available to Wren.
func newForeign(vm *C.WrenVM, x interface{}) {
	var (
		v   = reflect.Indirect(reflect.ValueOf(x))
		t   = v.Type()
		ptr = C.wrenSetSlotNewForeign(vm, C.int(0), C.int(0), C.size_t(t.Size()))
	)
	reflect.NewAt(t, ptr).Elem().Set(v)
}

//export wrengoResolveModule
func wrengoResolveModule(vm *C.WrenVM, importer *C.char, name *C.char) *C.char {
	path := C.CString(vmMap[vm].cb.ResolveModuleFunc(vmMap[vm], C.GoString(importer), C.GoString(name)))
	defer C.free(unsafe.Pointer(path))
	return path
}

//export wrengoLoadModule
func wrengoLoadModule(vm *C.WrenVM, name *C.char) *C.char {
	code := C.CString(vmMap[vm].cb.LoadModuleFunc(vmMap[vm], C.GoString(name)))
	defer C.free(unsafe.Pointer(code))
	return code
}

//export wrengoBindForeignMethod
func wrengoBindForeignMethod(vm *C.WrenVM, module *C.char, class *C.char, static C.bool, sign *C.char) unsafe.Pointer {
	m := C.GoString(module)
	if m != "main" {
		return unsafe.Pointer(nil)
	}

	var (
		className = C.GoString(class)
		isStatic  = bool(static)
		signature = C.GoString(sign)
	)

	if f, ok := vmMap[vm].methods[bindSignature(m, className, isStatic, signature)]; ok {
		return f
	}

	return unsafe.Pointer(nil)
}

//export wrengoBindForeignClass
func wrengoBindForeignClass(vm *C.WrenVM, module *C.char, className *C.char) C.WrenForeignClassMethods {
	m := C.GoString(module)
	if m != "main" {
		panic("tried to bind foreign class from non-main module")
	}

	cn := C.GoString(className)
	if c, ok := vmMap[vm].classes[cn]; ok {
		// Might be a good idea to support finalizers, but since this is Go,
		// I don't think they're actually necessary.
		return C.WrenForeignClassMethods{
			allocate: C.WrenForeignMethodFn(c),
			finalize: nil,
		}
	}

	panic(fmt.Sprintf("foreign class %s not found", cn))
}

//export wrengoWrite
func wrengoWrite(vm *C.WrenVM, text *C.char) {
	vmMap[vm].cb.WriteFunc(vmMap[vm], C.GoString(text))
}

//export wrengoError
func wrengoError(vm *C.WrenVM, err C.WrenErrorType, module *C.char, line C.int, message *C.char) {
	vmMap[vm].cb.ErrorFunc(vmMap[vm], ErrorType(err), C.GoString(module), int(line), C.GoString(message))
}

// Change 256 to a different number to enable more foreign class/method registrations.
//go:generate go run cgluer.go 256
