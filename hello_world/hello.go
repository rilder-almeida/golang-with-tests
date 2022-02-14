package main

import "fmt"

const (
	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func main() {
	fmt.Println(Hello("Rilder", "English"))
}

func Hello(name string, language string) string {
	switch language {
	case "Spanish":
		return spanishHelloPrefix + name
	case "French":
		return frenchHelloPrefix + name
	case "Portuguese":
		return portugueseHelloPrefix + name
	case "English":
		return englishHelloPrefix + name
	}
	if name == "" {
		return "Hello, nobody"
	}
	return "No greeting for you"
}
