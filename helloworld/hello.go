package main

import (
    "fmt"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) (helloMessage string) {
    helloMessage = englishHelloPrefix + name
    return helloMessage
}

func main() {
    fmt.Println(Hello("world"))
}

