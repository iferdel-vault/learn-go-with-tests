package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) (helloMessage string) {
	if name == "" {
		name = "World"
	}

    if language == spanish {
        return spanishHelloPrefix + name
    }
    if language == french {
        return frenchHelloPrefix + name
    }
	helloMessage = englishHelloPrefix + name
	return helloMessage
}

func main() {
	fmt.Println(Hello("world", ""))
}
