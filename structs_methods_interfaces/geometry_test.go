package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	// table driven tests
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 1.0, Height: 2.0}, hasArea: 2.0},
		{name: "Circle", shape: Circle{Radius: 2.0}, hasArea: 2.0 * 2.0 * math.Pi},
		{name: "Triangle", shape: Triangle{Base: 6.0, Heigth: 12.0}, hasArea: 36},
	}
	for _, tt := range areaTest {
		// using tt.name for case distinction between shapes in t.Run test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			hasArea := tt.hasArea
			assertCorrectComputation(t, tt.shape, got, hasArea)
		})
	}
}

func assertCorrectComputation(t testing.TB, shape Shape, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("%#v got '%0.2f' want '%0.2f'", shape, got, want)
	}
}
