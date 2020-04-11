package wrengo

import "fmt"

func defaultWrite(vm *VM, text string) {
	fmt.Print(text)
}

func defaultError(vm *VM, errorType ErrorType, module string, line int, message string) {
	fmt.Println(fmt.Sprint(errorType.String(), " (", module, ") Line ", line, " : ", message))
}
