package main

import (
	"testing"
)

func TestPlayground(t *testing.T) {
	CheckOutput(t, "playground.go", "Hello, World\n")
}
