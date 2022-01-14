package main

import "fmt"

func max(arrays []int, lenght int) int {
	for i := 0; i < lenght-1; i++ {
		if arrays[i] > arrays[i+1] {
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}
	maxValue := arrays[lenght-1]

	return maxValue
}

func main() {
	a := []int{12, 45, 90, 34, 12, 67}
	var lenght = len(a)

	fmt.Println(max(a, lenght))
}
