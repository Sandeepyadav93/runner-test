package main

import (
	"testing"
)

// comment
func TestAdd(t *testing.T) {
	want := 5
	got := Add(2, 3)
	if got != want {
		t.Fatalf("wanted %d got %d", want, got)
	}
}
