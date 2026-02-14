package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello go"
	got := hello()

	if want != got {
		t.Fatalf("expected %s, got %s\n", want, got)
	}
}
