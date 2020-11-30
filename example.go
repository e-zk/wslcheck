package example

import (
	"fmt"
	"github.com/e-zk/wslcheck"
	"log"
)

func main() {
	fmt.Printf("Am I running on WSL?\n")

	wsl, err := wslcheck.Check()
	if err != nil {
		log.Fatal(err)
	}

	if wsl == true {
		fmt.Printf("Yes!\n")
	} else {
		fmt.Printf("No!\n")
	}
}
