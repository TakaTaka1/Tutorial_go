package main

import (
	"fmt"
)

func SmallestIntegerFinder(numbers []int) int {
	// curr := numbers
	min := numbers[0]
	for _, v := range numbers {
		if min > v {
			min = v
		}
	}
	return min
}

func main() {

	nums := []int{100, 10, 100, 22, 4, 12, 13}
	fmt.Println(SmallestIntegerFinder(nums))
}
