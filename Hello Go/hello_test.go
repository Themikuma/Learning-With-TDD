package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Jorko")
	want := "Hello, Jorko"
	if got != want {

		t.Error(fmt.Printf("got %q want %q", got, want))
	}
}
