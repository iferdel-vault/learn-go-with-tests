package main

import (
	"fmt"
	"os"
	"regexp"
)

func CopyDigitsFromFile(filemame string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := os.ReadFile("./digits.txt")
	fmt.Println("after reading the digits.txt we got 'b': ", b)
	b = digitRegexp.Find(b)
	fmt.Println("after using Find method from regexp.MustCompile we have 'b': ", b)
	return append([]byte{}, b...)
}
