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

Works everything, but foreign part. And I doesn't check slots.

## Todo

- [ ] Create examples for slots
- [ ] Handle Foreign classes and methods
- [ ] Make intelligible binding
