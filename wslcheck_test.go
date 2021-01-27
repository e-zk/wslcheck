package wslcheck_test

import (
	"fmt"
	"github.com/e-zk/wslcheck"
	"log"
)

// The simplest possible use of this function
func ExampleCheck() {
	fmt.Printf("Am I running on WSL?\n")

	wsl, _ := wslcheck.Check()
	if wsl == true {
		fmt.Printf("Yes!\n")
	} else {
		fmt.Printf("No!\n")
	}
}
