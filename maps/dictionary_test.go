package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"go": "cool language"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("go")
		want := "cool language"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("gocha")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	t.Run("add word", func(t *testing.T) {
		word := "go"
		definition := "cool language"

		dictionary.Add(word, definition)
		got, err := dictionary.Search(word)

		assertError(t, err, nil)
		assertStrings(t, got, definition)
	})

	t.Run("add duplicate word", func(t *testing.T) {
		word := "test"
		definition := "added twice over"

		dictionary.Add(word, definition)
		err := dictionary.Add(word, definition)

		assertError(t, err, ErrDuplicateKey)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
