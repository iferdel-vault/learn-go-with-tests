package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("perform Sum with a fixed collection sized array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectSum(t, got, want, numbers)
	})
	t.Run("perform Sum with a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertCorrectSum(t, got, want, numbers)
	})
}

func assertCorrectSum(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d given collection of any size %v", got, want, numbers)
	}
}
