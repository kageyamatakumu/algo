package main

import (
	"bufio"
	"fmt"
	"os"
)

// 問題文
// パスワードを表す文字列 S が標準入力から与えられます。
// S が 6 文字以下の場合は dangerous、そうでない場合は safe を出力してください。

func main() {
	const passwordThreshold = 6

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	password := scanner.Text()

	if len(password) > passwordThreshold {
		fmt.Println("safe")
	} else {
		fmt.Println("dangerous")
	}
}