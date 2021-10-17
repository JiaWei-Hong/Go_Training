package main

import "fmt"

func main() {
	dict := make(map[string]string)

	dict["A"] = "B"
	dict["B"] = "C"
	dict["C"] = "D"
	dict["D"] = "E"

	for key, value := range dict {
		fmt.Printf("Key: %s Value: %s \n", key, value)
	}
}
