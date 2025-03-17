package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 2 つの正の整数 A,B が空白区切りで入力されます。A が B の倍数かどうかを判定してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 2 {
		fmt.Println("入力は2つの整数である必要があります。")
		return
	}

	num1, err := strconv.Atoi(textSlice[0])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	num2, err := strconv.Atoi(textSlice[1])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	if num1 <= 0 || num2 <= 0 {
		fmt.Println("入力は正の整数である必要があります。")
		return
	}

	if num1 % num2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}