package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 問題文
// 2 つの文字列 S,T が、標準入力で一行ずつ与えられます。
// これらの文字列をこの順に繋げて得られる文字列を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var builder strings.Builder

	for i := 0; i < 2; i++ {
		scanner.Scan()
		builder.WriteString(scanner.Text())
	}

	fmt.Println(builder.String())
}