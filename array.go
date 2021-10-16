package main

import "fmt"

func main() {
	gradeCount := 0
	grade := []int{}

	fmt.Println("請輸入要輸入幾筆資料:")
	fmt.Scanln(&gradeCount)

	for i := 0; i < gradeCount; i++ {
		score := 0

		fmt.Scanln(&score)

		grade = append(grade, score)

		fmt.Println("您輸入的成績為:", score)
	}

	sum := sum(grade)
	average := sum / len(grade)

	fmt.Println("總成績為:", sum, "平均為:", average)
}

func sum(scores []int) int {
	result := 0

	for _, v := range scores {
		result += v
	}

	return result
}
