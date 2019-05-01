package main

import "fmt"


const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"


func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	default:
		return englishPrefix
	}
}


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}


func main() {
	fmt.Println(Hello("world", ""))
}
