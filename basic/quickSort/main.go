package main

import "fmt"

func main() {
	data := []int{4, 2, 1, 5, 9, 6, 10}
	quickSort(data, 0, len(data)-1)
	fmt.Println(data)
}

func quickSort(data []int, first int, last int) {
	pivot := data[(first+last)/2] // data[3] = 5
	left := first                 // 0
	right := last                 // 6

	for {
		// left increment to pivot
		for data[left] < pivot {
			left++
		}
		// right decrement to pivot
		for data[right] > pivot {
			right--
		}

		if left >= right {
			break
		}

		tmp := data[left]
		data[left] = data[right]
		data[right] = tmp
		left++
		right--
	}
	// pivotの分があるからleft-1, right+1をしてから再帰処理
	if first < left-1 {
		quickSort(data, first, left-1)
	}

	if right+1 < last {
		quickSort(data, right+1, last)
	}

}
