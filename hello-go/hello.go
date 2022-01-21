package main

import "fmt"

const helloInEnglish = "Hello, "

// Hello writes out a greeting to the argument
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloInEnglish + name
}

func main() {
	fmt.Print(Hello("Jorko"))
}
