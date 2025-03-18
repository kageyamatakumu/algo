package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// 私たちが普段使っているグレゴリオ暦には、季節と暦のずれを補正するために平年より 1 日多いうるう年が存在します。
// うるう年は次の規則に従って定められています：
// 西暦年が 4 の倍数の年は (原則) うるう年である
// ただし、西暦年が 100 の倍数の年は (原則) うるう年でない
// ただし、西暦年が 400 の倍数の年は必ずうるう年である
// 例えば、 2020 年は 1. のみを満たすのでうるう年です。
// 西暦 N 年はうるう年でしょうか。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	year, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数で入力してください")
		return
	}

	// 西暦年が 400 の倍数の年は必ずうるう年である
	if year % 400 == 0 {
		fmt.Println("Yes")
		return
	}

	// 西暦年が 100 の倍数の年は (原則) うるう年でない
	if year % 100 == 0 {
		fmt.Println("No")
		return
	}

	// 西暦年が 4 の倍数の年は (原則) うるう年である
	if year%4 == 0 {
		fmt.Println("Yes")
		return
	}

	// それ以外は平年
	fmt.Println("No")
}