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
	// table driven tests
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{1.0, 2.0}, 2.0},
		{Circle{2.0}, 2.0 * 2.0 * math.Pi},
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		want := tt.want
		assertCorrectComputation(t, got, want)
	}
}

func assertCorrectComputation(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got '%0.2f' want '%0.2f'", got, want)
	}
}
