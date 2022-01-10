package main

import "fmt"

func main() {
	scores := []int{78, 90, 12, 45, 80}
	lenght := len(scores)

	insertSort(scores, lenght)
	for i := 0; i < lenght; i++ {
		fmt.Printf("%d,", scores[i])
	}

}

func insertSort(array []int, lenght int) {
	for i := 0; i < lenght; i++ {
		insertElement := array[i]
		insertPosition := i

		for j := insertPosition; j >= 0; j-- {
			if insertElement < array[j] {
				array[j+1] = array[j]
				insertPosition--
			}
		}

		array[insertPosition] = insertElement
	}
}
