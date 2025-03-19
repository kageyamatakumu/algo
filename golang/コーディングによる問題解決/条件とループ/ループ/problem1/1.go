package main

import (
	"fmt"
)

// 問題文
// 文字列 Hello を 10 回改行区切りで出力してください。

func main() {
	repeatStr := "Hello"

	for i := 0; i < 10; i++ {
		fmt.Println(repeatStr)
	}
}