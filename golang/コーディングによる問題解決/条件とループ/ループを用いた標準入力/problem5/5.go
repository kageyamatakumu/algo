package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// N 個の正の整数 S0,S1,…,SN-1が与えられます。N 個の文字列を前からすべてつなげてできる文字列の長さを出力してください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// 英小文字からなる長さ 1 以上 2000 以下の文字列 (0≤i≤N−1)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// N個の入力を受け取る
	scanner.Scan()
	inputN := scanner.Text()

	N, err := strconv.Atoi(inputN)
	if err != nil || N < 1 || N > 1000 {
		fmt.Println("1 以上 1000 以下の整数で入力してください")
		return
	}

	var strArr []string

	for i := 0; i < N; i++ {
		scanner.Scan()
		str := scanner.Text()

		// 各文字列の長さチェック
		if len(str) < 1 || len(str) > 2000 {
			fmt.Println("1 以上 2000 以下の文字列を入力してください")
			return
		}

		strArr = append(strArr, str)
	}

	// 文字列を結合し、その長さを出力
	fmt.Println(len(strings.Join(strArr, "")))
}