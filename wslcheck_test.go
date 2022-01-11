package wslcheck_test

import (
	"fmt"
	"go.zakaria.org/wslcheck"
	"testing"
)

// Test the version checking regexp
func TestCheckVer(t *testing.T) {
	var result bool

	result = wslcheck.CheckVer([]byte("4.2.0-microsoft-standard"))
	fmt.Println(result)

	result = wslcheck.CheckVer([]byte("4.4.0-19041-Microsoft"))
	fmt.Println(result)

	result = wslcheck.CheckVer([]byte("5.10.9-0-virt"))
	fmt.Println(result)

	// Output:
	// true
	// true
	// false
}

// The simplest possible use of this function
func ExampleCheck() {
	fmt.Println("Am I running on WSL?")

	wsl, _ := wslcheck.Check()
	if wsl == true {
		fmt.Println("Yes!")
	} else {
		fmt.Println("No!")
	}
}
