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

	helloMessage = greetingPrefix(language) + name

	return helloMessage
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case norsk:
		prefix = norskHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
