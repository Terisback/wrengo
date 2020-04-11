# Wrengo

Binding for Wren scripting language

## Usage

```go
package main

import (
    "github.com/Terisback/wrengo"
    "fmt"
)

func main() {
    config := wrengo.NewConfiguration()
    vm := wrengo.NewVM(config)
    fmt.Println(vm.Interpret("my_module", "System.print(\"I am running in a VM!\")"))
}
```

## Current

Only works creating a VM and Interpret

## Todo

- [ ] Make intelligible binding
