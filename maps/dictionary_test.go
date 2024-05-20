package main

import (
	"testing"
)

func TestDictionary(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		baseDefinition := "this is the original definition"
		dictionary := Dictionary{word: baseDefinition}

		updateDefinition := "this is an updated definition"
		dictionary.Update(word, updateDefinition)

		assertDefinition(t, dictionary, word, updateDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {

	t.Run("deleting existing word", func(t *testing.T) {
		word := "test"
		definition := "this is the original definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("expected %q to be deleted", word)
		}
	})
}
