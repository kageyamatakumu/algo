package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// カメのアルルはとあるテストで N 点を取りました。
// このテストの評定は、点数に応じて次のように決まります。
// 評定	テストの点数
// S	100 点
// A	90 点以上 99 点以下
// B	80 点以上 89 点以下
// C	70 点以上 79 点以下
// D	60 点以上 69 点以下
// E	50 点以上 59 点以下
// F	49 点以下
// アルルの点数を表すデータ N が標準入力から与えられます。
// アルルの評定を計算し出力するプログラムを作成してください。
// ただし、標準入力から誤って英小文字からなる文字列が与えられることもあります。
// そのような場合には invalid を出力してください。

func main() {
	// 0 点以上 100 点以下の範囲
	const validScoreMin = 0
	const validScoreMax = 100
	// 評定S
	const SScore = 100
	// 評定A
	const AScoreMin = 90
	const AScoreMax = 99
	// 評定B
	const BScoreMin = 80
	const BScoreMax = 89
	// 評定C
	const CScoreMin = 70
	const CScoreMax = 79
	// 評定D
	const DScoreMin = 60
	const DScoreMax = 69
	// 評定E
	const EScoreMin = 50
	const EScoreMax = 59

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	score, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("invalid")
		return
	}

	if score < validScoreMin || score > validScoreMax {
		fmt.Printf("%d 以上 %d 以下の範囲で入力してください\n", validScoreMin, validScoreMax)
		return
	}

	switch {
	case score == SScore:
		fmt.Println("S")
	case score >= AScoreMin && score <= AScoreMax:
		fmt.Println("A")
	case score >= BScoreMin && score <= BScoreMax:
		fmt.Println("B")
	case score >= CScoreMin && score <= CScoreMax:
		fmt.Println("C")
	case score >= DScoreMin && score <= DScoreMax:
		fmt.Println("D")
	case score >= EScoreMin && score <= EScoreMax:
		fmt.Println("E")
	default:
		fmt.Println("F")
	}

}