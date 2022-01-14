package main

import "fmt"

func swap(array []int, a, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}

func shellSorting(array []int, lenght int) {
	for gap := lenght / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < lenght; i++ {
			j := i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j = j - gap
			}
		}
	}
}

func main() {
	scores := []int{9, 6, 5, 8, 0, 7, 4, 3, 1, 2}
	lenght := len(scores)

	shellSorting(scores, lenght)

	for i := 0; i < lenght; i++ {
		fmt.Printf("%d,", scores[i])
	}

}
