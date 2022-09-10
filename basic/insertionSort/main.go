package main

import "fmt"

func main() {
	data := []int{5, 2, 14, 1, 3, 0}
	testInsertionSort(data)
	fmt.Println(data)
}

func testInsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		x := data[i]
		// i-1にすることでiの一つ前と比較することができる
		// data[j]>xが真の場合、data[j]とdata[j+1]=xを入れ替えることができる
		// jをデクリメントすることで、配列の内容に対して小さい値を左に1つずつ比較走査してソートできる
		for j := i - 1; j >= 0 && data[j] > x; j-- {

		}
	}
}

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		x := data[i]
		// data[1] x = 2
		// data[2] x = 14
		// data[3] x = 1
		// data[4] x = 3
		// data[5] x = 0
		for j := i - 1; j >= 0 && data[j] > x; j-- {
			// j=0, data[0]=5  {"2",5,14,1,3,0}
			// j=1, data[1]=5  {2,5,"14",1,3,0}=> data[j] > xが偽のためループは終了
			// j=2, data[2]=14 {2,5,"1",14,3,0}=> j=1 {2,"1",5,14,3,0}=> j=0 {"1",2,5,14,3,0}
			// j=3, data[3]=14 {1,2,5,"3",14,0}=> j=2 {1,2,"3",5,14,0}=> j=1以下は、x=3より残りの値が小さいため処理は実行されない
			// j=4, data[4]=14 {1,2,3,5,"0",14}=> j=3 {1,2,3,"0",5,14}=> j=2 {1,2,"0",3,5,14}=> j=1 {1,"0",2,3,5,14}=> j=0 {"0",1,2,3,5,14}
			tmp := data[j]
			data[j] = data[j+1]
			data[j+1] = tmp
		}
	}
}
