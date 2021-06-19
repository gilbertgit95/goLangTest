package main

import "testing"

func TestPlus(t *testing.T) {
	if Plus(1, 3) != 4 {
		t.Error("-> [1 + 3] expected answer is 4")
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(2, 4) != 8 {
		t.Error("-> [2 * 4] expected answer is 8")
	}
}
