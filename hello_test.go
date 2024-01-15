package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Clebs")
	want := "Hello, Clebs"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
