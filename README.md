# Wrengo

Binding for Wren scripting language

## Usage

```go
package main

import (
    "fmt"

    "github.com/Terisback/wrengo"
)

func main() {
    // Creating new VM
    config := wrengo.NewConfiguration()
    vm := wrengo.NewVM(config)
    defer vm.FreeVM()

    // Interpret "Hello World!" in Wren
    // It will print in console by default
    result := vm.Interpret("my_module", "System.print(\"Hello World!\")")

    // Will output RESULT_SUCCESS if it works
    fmt.Println(result.String())
}
```

## Current

Only works creating a VM and Interpret

## Todo

- [ ] Make intelligible binding
