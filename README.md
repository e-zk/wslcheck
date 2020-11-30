# wslcheck

_Am I running on WSL?_   
Go module to check whether your program is running on [Windows Subsystem for Linux](https://en.wikipedia.org/wiki/Windows_Subsystem_for_Linux) (WSL).

This module works by checking the contents of `/proc/sys/kernel/osrelease`.

## Usage

```console
go get github.com/e-zk/wslcheck
```

```go
package main

import (
	"fmt"
	"github.com/e-zk/wslcheck"
)

func main() {
	fmt.Printf("Am I running on WSL?\n")
	if wslcheck.Check() {
		fmt.Printf("Yes!\n")
	} else {
		fmt.Printf("No!\n")
	}
}
```