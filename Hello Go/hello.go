package main

import "fmt"

// Hello something
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Print(Hello("Jorko"))
}
