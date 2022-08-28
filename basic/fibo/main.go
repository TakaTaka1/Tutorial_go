package main

import "fmt"

func main() {

	// 0, 1, 1, 2, 3, 5, 8, 13, 21
	n := 4
	x, y := 0, 1

	fmt.Println(x)
	for i := 0; i < n; i++ {
		x, y = y, y+x
		fmt.Println(x)
	}
	_ = y

}
