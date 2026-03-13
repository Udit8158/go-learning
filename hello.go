package main

import "fmt"

func Hello(name string) string {

	if name == "" {
		return "Hello, World"
	}
	
	const englishHelloPrefix = "Hello"
	return  englishHelloPrefix + ", " + name
}

func Greet(name string) string {
	return "Hello, " + name
}


func main() {
	fmt.Println("We can do this too") // we don't need fmt import
}