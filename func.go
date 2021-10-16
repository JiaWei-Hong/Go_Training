package main

import "fmt"

func main() {
	var name string

	fmt.Print("Input your name:")
	fmt.Scanln(&name)

	hello(name)
}

func hello(name string) {
	fmt.Print("Nice to meet you " + name)
}
