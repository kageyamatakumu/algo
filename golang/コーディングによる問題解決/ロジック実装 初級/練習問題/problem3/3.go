package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// 標準入力から正の整数 N が与えられます。
// 税抜価格が N 円の商品を買った際の税込価格を求めてください。
// ただし、税率は 10% であるとします。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	if num < 0 {
		fmt.Println("入力は正の整数である必要があります。")
		return
	}

	price := num * 110 / 100

	fmt.Println(price)
}