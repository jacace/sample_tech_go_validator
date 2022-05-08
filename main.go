package main

import (
	"fmt"

	"github.com/jacace/platform"
)

func main() {
	fmt.Println("Validating platform capabitities...")
	if platform.IsOSSupported() {
		fmt.Println("Go compiler is present")
	}
}
