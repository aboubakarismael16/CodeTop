package main

import "fmt"

func min_value(array []int) int {
	minIndex := 0
	for i := 0; i < len(array)-1; i++ {
		if array[minIndex] > array[i] {
			minIndex = i
		}
	}

	return array[minIndex]
}

func main() {
	a := []int{12, 4, 5, 3, 100}

	minValue := min_value(a)
	fmt.Println("min Value is :", minValue)
}
