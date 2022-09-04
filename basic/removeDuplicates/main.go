package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 4, 5, 5, 6, 7, 8, 8, 8, 9}))
}

// refactoring
func removeDuplicates(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums[0:n]
	}
	j := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return nums[0:j]
}

// func removeDuplicates(nums []int) int {

// 	for i := 0; i < len(nums); i++ {
// 		if len(nums) == 1 {
// 			break
// 		}

// 		if len(nums) == 2 && nums[0] == nums[1] {
// 			nums = append(nums[:i], nums[i+1:]...)
// 			break
// 		}

// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				nums = append(nums[:i], nums[i+1:]...)
// 				i--
// 			}
// 			break
// 		}

// 	}
// 	return len(nums)
// }
