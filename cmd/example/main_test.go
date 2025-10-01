package main

import (
	"testing"
)

// TestMainRunsWithoutError verifies that main() executes without panicking
func TestMainRunsWithoutError(t *testing.T) {
	// This test ensures main() can run without errors
	// If main() panics or has runtime errors, this test will fail
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main() panicked: %v", r)
		}
	}()

	main()
}
