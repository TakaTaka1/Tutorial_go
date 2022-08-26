package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(11111))
}

func isPalindrome(x int) bool {
	checkStr := strconv.Itoa(x)
	flag := true
	length := len(checkStr)
	middle := length / 2 // 端数切り捨て

	for i := 0; i < middle; i++ {
		if checkStr[i] != checkStr[(length-1)-i] {
			return false
		}
	}

	return flag
}
