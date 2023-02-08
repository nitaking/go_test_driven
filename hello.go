package main

import "fmt"

const langSpanish = "Spanish"
const langFrench = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == langSpanish {
		return spanishHelloPrefix + name
	}

	if language == langFrench {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
