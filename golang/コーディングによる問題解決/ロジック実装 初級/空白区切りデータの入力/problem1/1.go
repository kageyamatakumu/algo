package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// 問題文
// 2 つの正の整数 A,B が空白区切りで入力されます。
// A+B の値を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 2 {
		fmt.Println("入力は2つの整数である必要があります。")
		return
	}

	var answer int = 0

	for _, numStr := range textSlice {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("入力は整数である必要があります。")
			return
		}
		answer += num
	}

	fmt.Println(answer)
}