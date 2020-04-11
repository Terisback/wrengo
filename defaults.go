package wrengo

import "fmt"

func defaultLoadModule(vm *VM, module string) string {
	fmt.Println("No LoadModuleFunc specified in configuration")
	return ""
}

// func defaultBindForeignMethod(vm *VM, module, className string, isStatic bool, signature string) {
// 	fmt.Println("BindForeignMethod in not implemented")
// }

func defaultBindForeignClass(vm *VM, module, className string) {
	fmt.Println("BindForeignClass in not implemented")
}

func defaultWrite(vm *VM, text string) {
	fmt.Print(text)
}

func defaultError(vm *VM, errorType ErrorType, module string, line int, message string) {
	fmt.Println(fmt.Sprint(errorType.String(), " (", module, ") Line ", line, " : ", message))
}
