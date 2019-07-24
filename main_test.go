package main

import "testing"

func TestSayHello(t *testing.T) {
	s := SayHello("Billy")
	if s != "Hello Billy!" {
		t.Errorf("string incorrect, got: %s, want: %s", s, "Hello Billy!")
	}
}
