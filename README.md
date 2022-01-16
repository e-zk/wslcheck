wslcheck [![GoDoc](https://godocs.io/github.com/e-zk/wslcheck?status.svg)](https://godocs.io/github.com/e-zk/wslcheck)
======== 

Go module to check whether your program is running on [Windows Subsystem for Linux](https://en.wikipedia.org/wiki/Windows_Subsystem_for_Linux).

This module works by checking the contents of `/proc/sys/kernel/osrelease`.

## Usage

```console
$ go get go.zakaria.org/wslcheck
```

```go
package main

import (
	"fmt"
	"github.com/e-zk/wslcheck"
)

func main() {
	fmt.Printf("Am I running on WSL?\n")

	wsl, _ := wslcheck.Check()
	if wsl == true {
		fmt.Printf("Yes!\n")
	} else {
		fmt.Printf("No!\n")
	}
}
```
