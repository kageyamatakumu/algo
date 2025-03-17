package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// 2 つの整数 a,b が、標準入力で一行ずつ与えられます。
// a+b の値を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0

	for i := 0; i < 2; i++ {
		scanner.Scan()
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("入力は整数である必要があります。")
			return
		}
		answer += num
	}

	fmt.Println(answer)
}