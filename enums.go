package wrengo

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

// The type of an object stored in a slot.
//
// This is not necessarily the object's *class*, but instead its low level
// representation type.
type WrenType int

const (
	WREN_TYPE_BOOL WrenType = iota
	WREN_TYPE_NUM
	WREN_TYPE_FOREIGN
	WREN_TYPE_LIST
	WREN_TYPE_NULL
	WREN_TYPE_STRING

	// The object is of a type that isn't accessible by the API.
	WREN_TYPE_UNKNOWN
)

func (i WrenType) String() string {
	return [...]string{
		"WREN_TYPE_BOOL",
		"WREN_TYPE_NUM",
		"WREN_TYPE_FOREIGN",
		"WREN_TYPE_LIST",
		"WREN_TYPE_NULL",
		"WREN_TYPE_STRING",
		"WREN_TYPE_UNKNOWN"}[i]
}
