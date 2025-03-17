package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// 0 以上 23 以下の整数 X が標準入力から与えられます。
// 現在の時刻が X 時のとき、日が変わる ( 24 時になる) まであと何時間かかるかを計算してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	if num < 0 || num > 23 {
		fmt.Println("0 以上 23 以下の整数を入力してください")
		return
	}

	fmt.Println(24 - num)
}