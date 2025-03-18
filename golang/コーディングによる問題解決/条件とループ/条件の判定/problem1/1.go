package main

import (
	"bufio"
	"fmt"
	"os"
)

// 問題文
// パスワードを表す文字列 S が標準入力から与えられます。
// S が password という文字列と一致している場合には dangerous、そうでない場合は safe を出力してください。

func main() {
	const dangerousPassword = "password"

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	password := scanner.Text()

	if password == dangerousPassword {
		fmt.Println("dangerous")
	} else {
		fmt.Println("safe")
	}
}