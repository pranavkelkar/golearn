package brego

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello()= %q, want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	a := 2
	b := 3
	want := 5
	if got := Sum(a, b); got != want {
		t.Errorf("Sum()= %q, want %q", got, want)
	}
}