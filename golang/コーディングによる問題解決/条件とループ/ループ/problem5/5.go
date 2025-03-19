package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// 次の 10 個の整数のうち、左から K 番目の整数を求めてください。
// 3,1,4,1,5,9,2,6,5,3
// ただし、最も左の整数を左から 0 番目と数えます
// K は 0 以上 9 以下の整数

func main() {
	CHECK_ARRAY := []int{3,1,4,1,5,9,2,6,5,3}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()


	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数を入力してください")
		return
	}

	if num < 0 || num > 9 {
		fmt.Println("0 以上 9 以下の整数を入力してください")
		return
	}

	fmt.Println(CHECK_ARRAY[num])

}