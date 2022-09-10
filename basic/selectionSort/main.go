package main

import "fmt"

func main() {
	data := []int{0, 1, 4, 10, 31, 10, 0, 2}
	selectionSort(data)
	fmt.Println(data)
}

// 仮の最小値のindexを求める(lowIndex)
// 求めたindexとネストしたループ内でlowIndex以降のindexの値を比較して、lowIndexより小さければ、
// lowIndexの値を更新する
// 外側ループのカウンタ変数i != lowIndexならば値の入れ替えを行う
func selectionSort(data []int) int {
	var indexLow int
	for i := 0; i < len(data)-1; i++ {
		indexLow = i
		for j := i + 1; j < len(data); j++ {
			if data[indexLow] > data[j] {
				indexLow = j
			}
		}
		if i != indexLow {
			tmp := data[i]
			data[i] = data[indexLow]
			data[indexLow] = tmp
		}
	}

	return indexLow
}
