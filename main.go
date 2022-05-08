package main

import (
	"fmt"

	"github.com/jacace/platform"
)

func main() {
	fmt.Println("Validating platform capabitities...")
	if platform.IsOSSupported("macOS") {
		fmt.Println("OS is supported")
	}
}
