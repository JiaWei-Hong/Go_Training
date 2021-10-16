package main

import "fmt"

func main() {
	x := 10

	add(x)

	fmt.Println("main function:", x)
}

func add(x int) {
	x += 10

	fmt.Println("add function:", x)
}
