package platform_test

import (
	"testing"

	"github.com/jacace/platform"
)

func TestIsOSSupported(t *testing.T) {
	if !platform.IsOSSupported("windows") {
		t.Fatal("OS is not supported")
	}
}
