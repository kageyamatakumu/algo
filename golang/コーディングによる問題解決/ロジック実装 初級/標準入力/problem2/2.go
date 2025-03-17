package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// 1 以上 100 以下の整数 N が標準入力から与えられます。
// N を 5 で割ったあまりを標準出力するプログラムを作成してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("入力は整数である必要があります")
		return
	}

	if num < 1 || num > 100 {
		fmt.Println("入力は1以上100以下の整数である必要があります。")
		return
	}

	fmt.Println(num % 5)
}