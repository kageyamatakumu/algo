package main

import (
	"bufio"
	"fmt"
	"os"
)

// 問題文
// 文字列 S が標準入力で与えられます。
// 文字列 S の先頭の文字を出力してください。 たとえば、文字列 algo の先頭の文字は a です。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputText := scanner.Text()

	fmt.Println(string([]rune(inputText)[0]))
	fmt.Println([]rune(inputText)[0])
	fmt.Println(inputText[0])
}