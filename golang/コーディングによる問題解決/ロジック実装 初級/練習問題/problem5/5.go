package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// 文字列 S と正の整数 N が改行区切りで入力されます。S の前から N 番目の文字を出力してください。
// ただしここでは、文字列 S の先頭の文字は 1 文字目であるとします。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()

	scanner.Scan()
	numStr := scanner.Text()
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("N は正の整数である必要があります。")
		return
	}

	if len(text) < num {
		fmt.Println("指定された位置の文字は存在しません。")
		return
	}

	fmt.Println(string(text[num-1]))
}