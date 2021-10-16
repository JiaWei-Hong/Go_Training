package main

import "fmt"

func main() {
	x := 10
	var xPtr *int = &x

	add(xPtr)

	fmt.Println("main function:", x)
}

func add(xPtr *int) {
	*xPtr += 10

	fmt.Println("add function:", *xPtr)
}
