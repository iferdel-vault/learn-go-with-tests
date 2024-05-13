package main

import (
    "fmt"
)

func Hello(name string) (helloMessage string) {
    helloMessage = "Hello, " + name
    return helloMessage
}

func main() {
    fmt.Println(Hello("world"))
}

