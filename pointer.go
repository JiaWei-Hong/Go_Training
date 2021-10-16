package main

import "fmt"

func main() {
	x := 5
	pointer := *&x

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(pointer)
}
