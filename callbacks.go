package wrengo

import "fmt"

func CallbackResolveModule(vm *VM, importer, name string) string {
	fmt.Println("CallbackLoadModule is not implemented")
	return ""
}

func CallbackLoadModule(vm *VM, module string) string {
	fmt.Println("CallbackLoadModule is not implemented")
	return ""
}

func CallbackBindForeignMethod(vm *VM, module, className string, isStatic bool, signature string) func(vm *VM) {
	fmt.Println("BindForeignMethod is not implemented")
	return func(vm *VM) {}
}

func CallbackBindForeignClass(vm *VM, module, className string) {
	fmt.Println("BindForeignClass is not implemented")
}

func CallbackWrite(vm *VM, text string) {
	fmt.Print(text)
}

func CallbackError(vm *VM, errorType ErrorType, module string, line int, message string) {
	fmt.Println(fmt.Sprint(errorType.String(), " (", module, ") Line ", line, " : ", message))
}
