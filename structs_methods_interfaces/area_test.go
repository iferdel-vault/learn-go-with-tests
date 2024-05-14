package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	got := Area(1.0, 2.0)
	want := 2.0

	if got != want {
		t.Errorf("got '%0.2f' want '%0.2f'", got, want)
	}
}
