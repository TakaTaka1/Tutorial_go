package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	var retInt []int
	for i := 0; i < len(nums); i++ {
		if len(retInt) == 2 {
			break
		}
		for j := 1 + i; j < len(nums); j++ {
			if target == (nums[i] + nums[j]) {
				retInt = append(retInt, i, j)
				break
			}
		}
	}

	return retInt
}

func main() {
	nums := []int{1, 2, 3, 45}
	target := 5
	fmt.Println(TwoSum(nums, target))
}
