package main

import "fmt"

type Item struct {
	name  string
	price int
}

func main() {
	apple := Item{"Apple", 10}
	banana := Item{"Banana", 20}

	fmt.Println(apple.name, apple.price)
	fmt.Println(banana.name, banana.price)
}
