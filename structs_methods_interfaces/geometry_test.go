package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{1.0, 2.0}
		got := rectangle.Perimeter()
		want := 6.0
		assertCorrectComputation(t, got, want)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{2.0}
		got := circle.Perimeter()
		want := 2.0 * math.Pi
		assertCorrectComputation(t, got, want)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectComputation(t, got, want)
	}

	t.Run("Rectangles", func(t *testing.T) {
		rectangle := Rectangle{1.0, 2.0}
		want := 2.0
		checkArea(t, rectangle, want)
	})
	t.Run("Circles", func(t *testing.T) {
		circle := Circle{2.0}
		want := 2.0 * 2.0 * math.Pi
		checkArea(t, circle, want)
	})
}

func assertCorrectComputation(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got '%0.2f' want '%0.2f'", got, want)
	}
}
