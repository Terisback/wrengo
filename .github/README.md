<p align="center">
    <a hrer="http://wren.io">
        <img alt="Wren" height="300" src="https://user-images.githubusercontent.com/26527529/79434497-eab31800-7fe7-11ea-842e-d86c9e46fdc0.png">
    </a>
    <br>
    <a href="https://pkg.go.dev/github.com/Terisback/wrengo?tab=doc">
        <img alt="go.dev" src="https://img.shields.io/badge/go.dev-007d9c?logo=go&logoColor=white&style=flat-square">
    </a>
    <a href="https://github.com/Terisback/wrengo/actions?query=workflow%3ABuild">
        <img alt="github actions build" src="https://img.shields.io/github/workflow/status/Terisback/wrengo/Build?style=flat-square&logo=github">
    </a>
</p>
<p align="center">
    <b>Wrengo</b> is binding for <b><a href="http://wren.io/">Wren</a></b> scripting language.
</p>

## ⚡️ Quickstart

```go
package main

import (
    "fmt"

    "github.com/Terisback/wrengo"
)

func main() {
    config := wrengo.NewConfiguration()
    config.WriteFunc = wrengo.CallbackWrite
    vm := wrengo.NewVM(config)
    defer vm.FreeVM()

    vm.Interpret("main", `System.print("Hello world!")`)
}
```

## ⚙️ Installation

First of all, [download](https://golang.org/dl/) and install Go. The project is written within `1.14`, but you can try your luck.

Installation is done using the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:

```bash
go get -u github.com/Terisback/wrengo
```

Also you need installed `gcc`:

On <b>Windows</b> you need [Mingw-w64](https://mingw-w64.org) or [TDM-GCC](http://tdm-gcc.tdragon.net/).
You can also build binary in [MSYS2](https://msys2.github.io/) shell.

On <b>MacOS</b> you need Xcode or Command Line Tools for Xcode.

## 🎯 Features

- <b>Wren is small.</b> The VM implementation is under [4,000 semicolons](https://github.com/wren-lang/wren/tree/master/src). You can skim the whole thing in an afternoon. It’s small, but not dense. It is readable and [lovingly-commented](https://github.com/wren-lang/wren/blob/46c1ba92492e9257aba6418403161072d640cb29/src/wren_value.h#L378-L433).
- <b>Wren is fast.</b> A fast single-pass compiler to tight bytecode, and a compact object representation help Wren [compete with other dynamic languages](http://wren.io/performance.html).
- <b>Wren is class-based.</b> There are lots of scripting languages out there, but many have unusual or non-existent object models. Wren places [classes](http://wren.io/classes.html) front and center.
- <b>Wren is concurrent.</b> Lightweight [fibers](http://wren.io/concurrency.html) are core to the execution model and let you organize your program into an army of communicating coroutines.

## 🏇🏻 Benchmark

| | fib(35) | fibt(35) |  Type  |
| :--- |    ---: |     ---: |  :---: |
| Go | `96ms` | `25ms` | Go (native) |
| [**Wrengo**](https://github.com/Terisback/wrengo) | `34ms` | `29ms` | Wren VM on Go |
| [Tengo](https://github.com/d5/tengo) | `34ms` | `30ms` | VM on Go |
| Lua | `2ms` | `25ms` | Lua (native) |
| [go-lua](https://github.com/Shopify/go-lua) | `7ms` | `26ms` | Lua VM on Go |
| [GopherLua](https://github.com/yuin/gopher-lua) | `7ms` | `29ms` | Lua VM on Go |
| Python | `5ms` | `58ms` | Python (native) |
| [starlark-go](https://github.com/google/starlark-go) | `16ms` | `26ms` | Python-like Interpreter on Go |
| [gpython](https://github.com/go-python/gpython) | `49ms` | `42ms` | Python Interpreter on Go |
| [goja](https://github.com/dop251/goja) | `8ms` | `28ms` | JS VM on Go |
| [otto](https://github.com/robertkrimen/otto) | `131ms` | `35ms` | JS Interpreter on Go |
| [Anko](https://github.com/mattn/anko) | `126ms` | `27ms` | Interpreter on Go |

_* [fib(35)](https://github.com/Terisback/wrengobench/blob/master/code/fib.wren):
Fibonacci(35)_  
_* [fibt(35)](https://github.com/Terisback/wrengobench/blob/master/code/fibtc.wren):
[tail-call](https://en.wikipedia.org/wiki/Tail_call) version of Fibonacci(35)_  
_* **Go** does not read the source code from file, while all other cases do_  
_* Results were rounded up_  
_* Tested on my Microsoft Sufrace 6 Pro_  
_* See [here](https://github.com/Terisback/wrengobench) for commands/codes used_

## 👨‍🦽 Compromises

<b>Foreign Function Limits</b> Due to Go's inability to generate C-exported functions at runtime, the number of foreign methods able to be registered with the Wren VM through this package is limited
to 256 for functions and 256 for classes. This number is completely arbitrary, though, and can be changed by modifying
the directive at the bottom of wrengo.go and running "go generate". If you feel like
this number is a terrible default, pull requests will be happily accepted.

## 👨🏻‍💻 Examples

Simplified examples is listed below. *(Without error checking and import)*

> Full examples are stored in the [cmd](https://github.com/Terisback/wrengo/tree/master/cmd) folder

### Interpret

```go
func main() {
	// New configuration for VM
	config := wrengo.NewConfiguration()

	// Adding callbacks
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError

	// Creating new VM
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	// New console read buffer
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ready! Press Ctrl+C for exit.")

	// Interpret loop
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		vm.Interpret("main", text+"\n")
	}
}
```

### Handles

```go
func main() {
	program := `
		class WrenMath {
			static do_add(a, b) {
				return a + b
			}
		}
	`

	config := wrengo.NewConfiguration()
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	vm.Interpret(wrengo.DefaultModule, program)

    // Make sure enough slots are allocated
    vm.EnsureSlots(3)

    // Getting class to slot 0
    vm.GetVariable(wrengo.DefaultModule, "WrenMath", 0)

    // Creating handle to do_add function
    h := vm.NewCallHandle("do_add(_,_)")

    // Setting call arguments
	vm.SetSlotDouble(1, 9)
    vm.SetSlotDouble(2, 3)

    // Calling
    h.Call()

    // Print the result
	fmt.Println(vm.GetSlotDouble(0))
}
```

### Foreign

```go
type God struct { msg string }

// The constructor for a foreign type takes no arguments and returns
// an interface{} value representing the new object.
func NewGod() interface{} {
	return &God{msg: "What are you doing? %s"}
}

// Function to bind into Wren
func GetGodsMessage(vm *wrengo.VM) {
    // Getting foreign class
    god := vm.GetSlotForeign(0, God{}).(God)

    // Getting argument
    name := vm.GetSlotString(1)

    // Return result
	vm.SetSlotString(0, fmt.Sprintf(god.msg, name))
}

func main() {
	// Wren code
	program := `
			foreign class God {
				construct new() {}
				foreign getMessage(name)
			}

			var god = God.new()
			System.print(god.getMessage("Silly boy"))
		`

	config := wrengo.NewConfiguration()
	config.WriteFunc = wrengo.CallbackWrite
	config.ErrorFunc = wrengo.CallbackError
	vm := wrengo.NewVM(config)
	defer vm.FreeVM()

	// Bind God class
	vm.BindForeignClass("God", NewGod)

	// Bind God.getMessage()
	vm.BindForeignMethod("God", false, "getMessage(_)", GetGodsMessage)

	vm.Interpret(wrengo.DefaultModule, program)
}
```

## 📚 Similar/nearby projects

- [go-wren](https://github.com/dradtke/go-wren) - it's based on nearly 0.1.0 Wren sources, when Wrengo based on 0.2.0 Wren release. It's doing same work as this package, but working with older version of Wren and it has more simplifications.
- [tengo](https://github.com/d5/tengo) - Tengo is fast and secure because it's compiled/executed as bytecode on stack-based VM that's written in native Go.
- [go-lua](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go.

More in [awesome-go](https://github.com/avelino/awesome-go#embeddable-scripting-languages) repository.
