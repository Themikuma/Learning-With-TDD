package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Error(fmt.Printf("got %q want %q", got, want))
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Jorko", "")
		want := "Hello, Jorko"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when emprty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Javier", spanish)
		want := "Hola, Javier"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Pier", french)
		want := "Bonjour, Pier"
		assertCorrectMessage(t, got, want)
	})
}
