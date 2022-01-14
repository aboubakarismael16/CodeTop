package main

import "fmt"

func reverse(array []int, lenght int) {
	middle := lenght / 2
	for i := 0; i <= middle; i++ {
		temp := array[i]
		array[i] = array[lenght-i-1]
		array[lenght-i-1] = temp
	}
}

func main() {
	scores := []int{50, 60, 70, 80, 90}
	lenght := len(scores)

	reverse(scores, lenght)
	for i := 0; i < lenght; i++ {
		fmt.Printf("%d,", scores[i])
	}

}
