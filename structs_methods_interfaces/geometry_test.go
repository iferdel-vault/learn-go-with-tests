package main

import (
	"testing"
)

func TestRectanglePerimeter(t *testing.T) {
    rectangle := Rectangle(1.0, 2.0)
    got := Perimeter(rectangle)
	want := 6.0

	if got != want {
		t.Errorf("got '%0.2f' want '%0.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
    rectangle := Rectangle(1.0, 2.0)
	got := Area(rectangle)
	want := 2.0

	if got != want {
		t.Errorf("got '%0.2f' want '%0.2f'", got, want)
	}
}

