package main

import "fmt"

func permutations(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	} else {
		var result = [][]int{}
		var input = []int{}
		var rest = [][]int{}
		for i, v := range nums {
			input = append(input, nums[:i]...)
			input = append(input, nums[i+1:]...)

			rest = append(rest, permutations(input)...)
			input = nil
			for _, restRange := range rest {
				val := []int{v}
				val = append(val, restRange...)
				result = append(result, val)
			}
			rest = nil
		}
		return result
	}
}

func main() {
	test := []int{2, 0, 1}
	fmt.Println(permutations(test))
	// [[2 0 1] [2 1 0] [0 2 1] [0 1 2] [1 2 0] [1 0 2]]
}
