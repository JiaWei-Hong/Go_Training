package main

import "fmt"

func main() {
	a, b := 0, 0

	fmt.Println("請輸入兩個數字中間使用空白隔開:")
	fmt.Scanln(&a, &b)
	fmt.Println(a + b)
}
