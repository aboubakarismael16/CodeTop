package main

import "fmt"

func main() {
	scores := []int{50, 68, 46, 70, 90, 80}
	var value int
	fmt.Printf("Entry the value you find for: ")
	fmt.Scan(&value)

	lenght := len(scores)
	isValue := false
	for i := 0; i < lenght; i++ {
		if scores[i] == value {
			isValue = true
			fmt.Printf("The value is %d, the index is %d ", value, i)
			break
		}
	}

	if !isValue {
		fmt.Printf("the value %d is not found ", value)
	}

}
