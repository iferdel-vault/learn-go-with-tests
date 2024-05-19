package main

import (
	"fmt"
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
    t.Run("new word", func (t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is a test"

	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
    })

    t.Run("existing word", func (t *testing.T) {
	word := "test"
	definition := "this is a test"
    dictionary := Dictionary{word: definition}
    fmt.Println("printing dictionary with data", dictionary)
        

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
