package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// カメのアルルがとあるテストを受けたところ、点数は N 点でした。
// 正整数 N が標準入力から与えられます。
// テストの点数が 0 点以上 100 点以下であるならば valid、そうでないならば invalid を出力してください。



func main() {
	const validScoreMin = 0
	const validScoreMax = 100

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	score, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数を入力してください")
		return
	}

	if score >= validScoreMin && score <= validScoreMax {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}

}