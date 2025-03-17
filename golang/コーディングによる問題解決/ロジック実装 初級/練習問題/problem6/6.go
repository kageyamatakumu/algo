package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 4 つの正の整数 A,B,C,D が空白区切りで入力されます。4 つの整数の最大値を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())
	if len(textSlice) != 4 {
		fmt.Println("入力は4つの整数である必要があります。")
		return
	}

	var answer int

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