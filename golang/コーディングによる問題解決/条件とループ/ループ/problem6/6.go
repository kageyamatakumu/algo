package main

import "fmt"

// 問題文
// 繰り返し構文を活用して、次の 10 個の整数の総和を計算し出力するプログラムを作成してください。
// 3, 1, 4, 1, 5, 9, 2, 6, 5, 3

func main() {
	arr := [10]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}

	var sum int

	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)
}