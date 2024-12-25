package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Clebson")
	want := "Hello, Clebson"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
