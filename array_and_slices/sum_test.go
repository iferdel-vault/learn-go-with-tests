package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
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

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Perform sum of tails across multiple slices saving each one into new slice",
		func(t *testing.T) {
			got := SumAllTails([]int{1, 3}, []int{4, 4, 2})
			want := []int{3, 6}
			checkSums(t, got, want)
		})

	t.Run("safely sum empty slices",
		func(t *testing.T) {
			got := SumAllTails([]int{}, []int{4, 4, 2})
			want := []int{0, 6}
			checkSums(t, got, want)
		})
}
