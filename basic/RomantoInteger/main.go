package main

import "fmt"

func main() {
	fmt.Println(RomanToInteger("MCMC")) // 1900
}

func RomanToInteger(s string) int {
	var roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	currentNum, lastNum, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		currentNum = roman[char]
		if currentNum < lastNum {
			total = total - currentNum
		} else {
			total = total + currentNum
		}

		lastNum = currentNum
	}
	return total
}
