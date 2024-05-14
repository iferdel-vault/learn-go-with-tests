package main

import (
    "testing"
)

func TestPerimeter(t *testing.T) {
    got := Perimeter(1, 2)
    want := 6

    if got != want {
        t.Errorf("got '%d' want '%d'", got, want)
    }
}

