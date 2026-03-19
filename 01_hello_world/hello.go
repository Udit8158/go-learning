package hello

import (
	"fmt"
)

const (
	spanishHelloPrefix = "Hola"
	englishHelloPrefix = "Hello"
	frenchHellowPrefix = "Bonjur"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	default:
		prefix = englishHelloPrefix + ", "

	case "Spanish":
		prefix = spanishHelloPrefix + ", "

	case "French":
		prefix = frenchHellowPrefix + ", "

	}

	return
}

func Greet(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println("We can do this too") // we don't need fmt import
}
