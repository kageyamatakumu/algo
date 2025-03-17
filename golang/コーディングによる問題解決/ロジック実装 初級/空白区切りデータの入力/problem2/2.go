package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 2 つの正の整数 A,B が空白区切りで入力されます。 A と B のうち大きい方の値を出力してください。
// ただし、A と B の値は異なることが保証されています。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 2 {
		fmt.Println("入力は2つの整数である必要があります。")
		return
	}

	answer := 0

	for _, numStr := range textSlice {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("入力は整数である必要があります。")
			return
		}

		if answer < num {
			answer = num
		}
	}

	fmt.Println(answer)
}