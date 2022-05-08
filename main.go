package main

import (
	"fmt"

	"github.com/jacace/platform"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Validating platform capabitities...")
	var platformName = "macOS"
	var isSupported, error = platform.IsOSSupported("macOS")
	if error == nil && isSupported {
		fmt.Printf("OS %v is supported", platformName)
		fmt.Println(quote.Go())
	}
}
