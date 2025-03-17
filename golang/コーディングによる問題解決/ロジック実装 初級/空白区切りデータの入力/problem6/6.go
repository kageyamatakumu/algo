package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

// 問題文
// 3 つの文字列 S, T, U が空白区切りで入力されます。U と T と S をこの順につなげた文字列を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 3 {
		fmt.Println("空白区切りで3つの文字列を入力してください。")
		return
	}

	var builder strings.Builder

	builder.WriteString(textSlice[2])
	builder.WriteString(textSlice[1])
	builder.WriteString(textSlice[0])

	fmt.Println(builder.String())
}