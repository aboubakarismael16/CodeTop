package main

import "fmt"

func main() {
	scores := []int{90, 70, 50, 80, 60, 85}
	lenght := len(scores)

	temp := make([]int, lenght+1)

	insert(scores, temp, lenght, 75, 3)

	scores = temp
	for i := 0; i < lenght+1; i++ {
		fmt.Printf("%d,", scores[i])
	}

}

func insert(score, temp []int, lenght, value, index int) {
	for i := 0; i < lenght; i++ {
		if i < index {
			temp[i] = score[i]
		} else {
			temp[i+1] = score[i]
		}
	}
	temp[index] = value
}
