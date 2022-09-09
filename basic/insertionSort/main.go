package main

import "fmt"

func main() {
	// Your code here!
	data := []int{5, 2, 14, 1, 3, 0}
	fmt.Println(insertionSort(data))
}

func insertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		x := data[i]
		for j := i - 1; j >= 0 && data[j] > x; j-- {
			tmp := data[j]
			data[j] = data[j+1]
			data[j+1] = tmp
		}
	}
	return data
}
