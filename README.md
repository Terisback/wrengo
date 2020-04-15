<p align="center">
    <a hrer="http://wren.io">
        <img alt="Wren" height="125" src="http://wren.io/wren.svg">
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
    <b>Wrengo</b> is binding for <b>Wren</b> scripting language.
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

    err := vm.Interpret("main", `System.print("Hello world!")`)
    if err != nil {
        fmt.Println(err)
    }
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

## 👀 Examples

A few examples are stored in the [cmd](https://github.com/Terisback/wrengo/tree/master/cmd) folder

## 🧬 Future plans

- [ ] Implement Foreign classes and methods
- [ ] Make intelligible binding
