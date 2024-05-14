package main

import (
	"os"
	"regexp"
)

func CopyDigitsFromFile(filemame string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := os.ReadFile("./digits.txt")
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
