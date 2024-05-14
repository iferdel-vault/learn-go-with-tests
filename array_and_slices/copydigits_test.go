package main

import (
	"slices"
	"testing"
)

func TestCopyDigitsFromFile(t *testing.T) {
	got := CopyDigitsFromFile("./digits.txt")
	want := []byte{56, 56, 57}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
