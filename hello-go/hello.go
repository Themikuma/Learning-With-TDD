package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloInEnglish = "Hello, "
const helloInSpanish = "Hola, "
const helloInFrench = "Bonjour, "

// Hello writes out a greeting to the argument
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getLanguagePrefix(language) + name
}

func getLanguagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloInSpanish
	case french:
		prefix = helloInFrench
	default:
		prefix = helloInEnglish
	}
	return prefix
}

func main() {
	fmt.Print(Hello("Jorko", ""))
}
