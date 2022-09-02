package main

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(test)
	// removeDuplicates([]int{0, 0, 1, 1, 1})
	// test = append(test[:2], test[3:]...)
	// fmt.Println(test[:3]) // index 3未満
	// fmt.Println(test[3:]) // index 3から
}

func removeDuplicates(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if len(nums) == 1 || len(nums) == 2 {
			break
		}

		if len(nums) == 2 && nums[0] == nums[1] {
			nums = append(nums[:i], nums[i+1:]...)
			break
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
			break
		}

	}
	return len(nums)
}
