package fib

import "testing"

func TestFibBaseCases(t *testing.T) {
	if res := Fib(0); res != 0 {
		t.Errorf("fib(0) != %d", res)
	}
	if res := Fib(1); res != 1 {
		t.Errorf("fib(1) != %d", res)
	}
}

func TestFib(t *testing.T) {
	if res := Fib(20); res != 6765 {
		t.Errorf("fib(20) != %d", res)
	}
}
