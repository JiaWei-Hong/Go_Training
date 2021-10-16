package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i >= j {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}

		fmt.Println()
	}
}
