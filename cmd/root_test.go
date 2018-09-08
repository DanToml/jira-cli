package cmd

import (
	"testing"
)

func TestRootCmd__HasShortDesc(t *testing.T) {
	if rootCmd.Short == "" {
		t.Fatal("Expected rootCmd.Short to be non-empty")
	}
}

func TestRootCmd__HasLongDesc(t *testing.T) {
	if rootCmd.Long == "" {
		t.Fatal("Expected rootCmd.Long to be non-empty")
	}
}
