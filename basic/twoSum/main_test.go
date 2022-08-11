package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testArr := []struct {
		inputArr        []int
		expectationNums []int
		target          int
	}{
		{
			inputArr:        []int{1, 2, 3, 4, 5},
			expectationNums: []int{2, 4},
			target:          8,
		},
		{
			inputArr:        []int{1, 2, 3, 4, 5},
			expectationNums: []int{3, 4},
			target:          9,
		},
	}

	for _, v := range testArr {
		t.Run("TwoSum", func(t *testing.T) {
			result := TwoSum(v.inputArr, v.target)
			if !reflect.DeepEqual(result, v.expectationNums) {
				fmt.Println("Not equal")
			}
		})
	}
}
