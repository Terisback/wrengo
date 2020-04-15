package wrengo

import "fmt"

// Default write callback that outputs to Stdout
func CallbackWrite(vm *VM, text string) {
	fmt.Print(text)
}

// Default error callback that outputs to Stdout
func CallbackError(vm *VM, errorType ErrorType, module string, line int, message string) {
	fmt.Println(fmt.Sprint(errorType.String(), " (", module, ") Line ", line, " : ", message))
}
