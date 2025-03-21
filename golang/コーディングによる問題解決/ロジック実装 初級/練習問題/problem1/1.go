package main

import (
	"bufio"
	"fmt"
	"os"
)

// 問題文
// 長さ 5 の文字列 S が標準入力から与えられます。文字列 S の中央の文字を出力してください。
// たとえば入出力例 1 に示す通り、S= power に対しては、文字 w を出力します。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	text := scanner.Text()
	if len(text) != 5 {
		fmt.Println("長さ 5 の文字を入力してください")
		return
	}

	fmt.Println(string(text[2]))
}