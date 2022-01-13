package main

import "fmt"

func binarySearch(array []int, lenght, searchValue int) int {
	low := 0
	high := lenght - 1
	mid := 0

	for {
		if low >= high {
			break
		}

		mid = (low + high) / 2
		if array[mid] == searchValue {
			return mid
		} else if array[mid] < searchValue {
			low = mid + 1
		} else if array[mid] > searchValue {
			high = mid - 1
		}

	}

	return -1
}

func main() {
	scores := []int{30, 40, 50, 60, 70, 80, 90}
	lenght := len(scores)
	searchValue := 40

	position := binarySearch(scores, lenght, searchValue)
	fmt.Printf("%d postion: %d", searchValue, position)

}
