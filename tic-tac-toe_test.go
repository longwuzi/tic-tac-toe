package main

import "testing"

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("Test fail expected: %s, result: %s\n", "o", b.get(1, 1))
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "x")
	if b.get(1, 1) != "x" {
		t.Errorf("Test fail expected: %s, result: %s\n", "x", b.get(1, 1))
	}
}
