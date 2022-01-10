package main

import (
	"fmt"
)

func main() {
	scores := []int{76, 86, 48, 90}
	lenght := len(scores)

	temp := make([]int, lenght+1)
	for i := 0; i < lenght; i++ {
		temp[i] = scores[i]
	}
	temp[lenght] = 75

	scores = temp
	for i := 0; i < lenght+1; i++ {
		fmt.Printf("%d,", scores[i])
	}

}
