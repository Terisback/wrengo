// +build ignore

// This file is intended to be executed by "go generate".

package main

import (
	"os"
	"strconv"
	"text/template"
)

var fileFuncTemplate = template.Must(template.New("").Parse(`package wrengo

/*
#include "wren.h"
{{range .}}extern void f{{.}}(WrenVM*);
{{end}}
static inline void* get_f(int i) {
	switch (i) {
		{{range .}}case {{.}}: return f{{.}};
		{{end}}default: return (void*)(0);
	}
}
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

const MAX_FUNC_REGISTRATIONS = {{len .}}

var (
	fMap = make(map[int]func(*VM))
	fMapGuard sync.Mutex
	fCounter int
)

{{range .}}
//export f{{.}}
func f{{.}}(vm *C.WrenVM) {
	f := fMap[{{.}}]
	if f == nil {
		panic("function {{.}} not registered")
	}
	f(vmMap[vm])
}
{{end}}

func registerFunc(name string, f func(*VM)) (unsafe.Pointer, error) {
	if (fCounter+1) >= MAX_FUNC_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	fMapGuard.Lock()
	defer fMapGuard.Unlock()

	fMap[fCounter] = f
	ptr := C.get_f(C.int(fCounter))
	fCounter++
	return ptr, nil
}
`))

var fileClassTemplate = template.Must(template.New("").Parse(`package wrengo

/*
#include "wren.h"
{{range .}}extern void c{{.}}(void *vm);
{{end}}
static inline void* get_c(int i) {
	switch (i) {
		{{range .}}case {{.}}: return c{{.}};
		{{end}}default: return (void*)(0);
	}
}
*/
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

const MAX_CLASS_REGISTRATIONS = {{len .}}

var (
	cMap = make(map[int]func())
	cMapGuard sync.Mutex
	cCounter int
)

{{range .}}
//export c{{.}}
func c{{.}}(vm unsafe.Pointer) {
	f := cMap[{{.}}]
	if f == nil {
		panic("function {{.}} not registered")
	}
	f()
}
{{end}}

func registerClass(name string, f func()) (unsafe.Pointer, error) {
	if (cCounter+1) >= MAX_CLASS_REGISTRATIONS {
		return nil, errors.New("maximum function registration reached")
	}

	cMapGuard.Lock()
	defer cMapGuard.Unlock()

	cMap[cCounter] = f
	ptr := C.get_c(C.int(cCounter))
	cCounter++
	return ptr, nil
}
`))

func main() {
	if len(os.Args) == 1 {
		panic("no number provided")
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	f, err := os.Create("funcGlue.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	if err := fileFuncTemplate.Execute(f, data); err != nil {
		panic(err)
	}

	c, err := os.Create("classGlue.go")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	data = make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	if err := fileClassTemplate.Execute(c, data); err != nil {
		panic(err)
	}
}
