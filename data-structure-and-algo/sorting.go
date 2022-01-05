package main

import "fmt"

func sort(array []int, lenght int) {
	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-i-1; j++ {
			if array[j] > array[j+1] {
				var flag = array[j]
				array[j] = array[j+1]
				array[j+1] = flag
			}
		}
	}
}

func main() {
	scores := []int{58, 80, 68, 70, 95, 56}
	var lenght = len(scores)

	sort(scores, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d,", scores[i])
	}
}
