package main

import "testing"

func testPlus(t *testing.T) {
	if Plus(1, 3) != 4 {
		t.Error("-> expected answer is 4")
	}
	if Plus(1, 2) != 4 {
		t.Error("-> expected answer is 4")
	}
}
