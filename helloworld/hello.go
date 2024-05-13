package main

import (
	"fmt"
)

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) (helloMessage string) {
	if name == "" {
		name = "World"
	}

    if language == spanish {
        return spanishHelloPrefix + name
    }
	helloMessage = englishHelloPrefix + name
	return helloMessage
}

func main() {
	fmt.Println(Hello("world", ""))
}
