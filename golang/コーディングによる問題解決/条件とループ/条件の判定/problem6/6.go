package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// カメのアルルは数学のテストで N 点、英語のテストで M 点を取りました。
// 数学と英語がともに 70 点以上 かつ、数学と英語の合計点が 160 点以上であれば試験に合格です。
// 標準入力から正整数 N, M の値が与えられます。
// アルルが試験に合格したならば Yes、そうでないならば No を出力するプログラムを作成してください。

func main() {
	 // 各科目の最低合格点
	const subjectPassScore = 70
	// 合計点の最低合格点
	const totalPassScore = 160
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	sliceString := strings.Fields(scanner.Text())

	if len(sliceString) < 2 {
		fmt.Println("入力形式が不正です")
		return
	}

	scoreOfMath, err := strconv.Atoi(sliceString[0])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	if scoreOfMath < subjectPassScore {
		fmt.Println("No")
		return
	}

	scoreOfEnglish, err := strconv.Atoi(sliceString[1])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	if scoreOfEnglish < subjectPassScore {
		fmt.Println("No")
		return
	}

	if scoreOfMath + scoreOfEnglish < totalPassScore {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}


}