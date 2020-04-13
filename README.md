# Wrengo

Binding for Wren scripting language

## Usage

```go
package main

import (
    "bufio"
    "fmt"
    "os"

    "github.com/Terisback/wrengo"
)

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

## Current

Works everything, but foreign part. And I doesn't check slots.

## Todo

- [ ] Create examples for slots
- [ ] Handle Foreign classes and methods
- [ ] Make intelligible binding
