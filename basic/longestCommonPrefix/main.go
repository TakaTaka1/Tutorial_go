package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(s []string) string {
	var prefix string = ""

	if len(s) > 0 {
		sort.Strings(s)
		first := string(s[0])
		last := string(s[len(s)-1])

		for i := 0; i < len(first); i++ {
			if first[i] == last[i] {
				prefix += string(first[i])
			} else {
				prefix = "Not matched"
				break
			}
		}
	}
	return prefix
}

func main() {
	test1 := []string{"alert", "alarm", "low"}
	fmt.Println(longestCommonPrefix(test1))
	test2 := []string{"win", "window", "windows"}
	fmt.Println(longestCommonPrefix(test2))
}
