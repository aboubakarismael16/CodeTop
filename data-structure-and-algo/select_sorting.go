package main

import "fmt"

func select_sort(array []int, lenght int) {
	var minIndex int

	for i := 0; i < lenght-1; i++ {
		minIndex = i
		minValue := array[minIndex]
		for j := i; j < lenght-1; j++ {
			if minValue > array[j+1] {
				minValue = array[j+1]
				minIndex = j + 1
			}
		}

		if i != minIndex {
			temp := array[i]
			array[i] = array[minIndex]
			array[minIndex] = temp
		}
	}
}

func main() {
	a := []int{45, 69, 46, 23, 90, 88}
	lenght := len(a)

	select_sort(a, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf(" %d,", a[i])
	}

}
