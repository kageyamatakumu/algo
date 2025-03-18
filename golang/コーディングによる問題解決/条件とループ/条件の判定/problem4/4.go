package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// カメのアルルはとあるテストで N 点を取りました。
// このテストの評定は、点数に応じて次のように決まります。
// 評定	テストの点数
// A	90 点以上 100 点以下
// B	80 点以上 89 点以下
// C	79 点以下
// 正整数 N が標準入力から与えられます。
// アルルの評定を計算し出力するプログラムを作成してください。

func main() {
	const validScoreMin = 0
	const validScoreMax = 100
	const AScoreMin = 90
	const AScoreMax = 100
	const BScoreMin = 80
	const BScoreMax = 89

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	score, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数を入力してください")
		return
	}

	if score < validScoreMin || score > validScoreMax {
		fmt.Println("0 点以上 100 点以下の範囲で入力してください")
		return
	}

	if score <= AScoreMax && score >= AScoreMin {
		fmt.Println("A")
	} else if score <= BScoreMax && score >= BScoreMin {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}