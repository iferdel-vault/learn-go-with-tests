package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"
	norsk   = "Norsk"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	norskHelloPrefix   = "Hallaisen, "
)

func Hello(name string, language string) (helloMessage string) {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case norsk:
		prefix = norskHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
