package main

import "fmt"

func main() {

	fmt.Println(LowestAndLargest(CheckPrime(1, 100)))

}

func CheckPrime(lowest, largest int) []int {
	primeList := []int{}
	for i := lowest; i <= largest; i++ {
		if i == 0 || i == 1 {
			continue
		}

		flag := 1

		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				flag = 0
				break
			}
		}

		if flag == 1 {
			primeList = append(primeList, i)
		}
	}
	return primeList
}

func LowestAndLargest(primeList []int) map[string]int {
	retList := make(map[string]int)

	lowest, largest := primeList[0], primeList[0]
	for i := 0; i < len(primeList); i++ {
		if primeList[i] > largest {
			largest = primeList[i]
		} else {
			lowest = primeList[i]
		}
	}

	retList["lowest"] = lowest
	retList["largest"] = largest

	return retList
}
