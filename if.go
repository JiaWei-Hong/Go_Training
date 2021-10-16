package main

import "fmt"

func main() {
	money := 0

	fmt.Println("Take out money?")
	fmt.Scanln(&money)

	if money < 100 {
		fmt.Println("Too Few")
	} else if money <= 20000 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Too Much")
	}

	fmt.Println("Done...")
}
