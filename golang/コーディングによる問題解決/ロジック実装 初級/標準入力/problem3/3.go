package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 文字列 S が標準入力で与えられます。
// S を 3 回繰り返してできる文字列を出力するプログラムを作成してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputText := scanner.Text()

	fmt.Println(strings.Repeat(inputText, 3))
}