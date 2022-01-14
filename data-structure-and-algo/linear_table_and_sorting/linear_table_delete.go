package main

import "fmt"

func main() {
	scores := []int{90, 70, 75, 80, 50, 45}
	fmt.Printf("Enter the index of value you need to delete:")
	var index int
	fmt.Scan(&index)

	lenght := len(scores)
	temp := make([]int, lenght-1)

	for i := 0; i < lenght; i++ {
		if i < index {
			temp[i] = scores[i]
		}
		if i > index {
			temp[i-1] = scores[i]
		}
	}

	scores = temp
	for i := 0; i < lenght-1; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
