package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

// 問題文
// 2 つの文字列 S,T が空白区切りで入力されます。S と T が等しい文字列であるかを判定してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 2 {
		fmt.Println("入力は2つである必要があります。")
		return
	}

	if textSlice[0] == textSlice[1] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}