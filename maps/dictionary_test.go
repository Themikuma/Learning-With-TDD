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

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "a measure of correctness"
		newDefinition := "a procedure intended to establish the quality of a product"

		dictionary.Add(word, definition)
		dictionary.Update(word, newDefinition)

		got, _ := dictionary.Search(word)

		assertStrings(t, got, newDefinition)
	})

	t.Run("update non existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "a procedure intended to establish the quality of a product"

		dictionary.Update(word, newDefinition)

		_, got := dictionary.Search(word)

		assertError(t, got, ErrNotFound)
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
