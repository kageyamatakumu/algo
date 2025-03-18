package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// ある大学の入学試験には 1 次試験と 2 次試験があり、受験者の合否は 1 次試験と 2 次試験の点数に応じて次のように決まります。
// 1 次試験の点数が X 点未満であれば、不合格
// 上記以外の場合、1 次試験と 2 次試験の合計点が Y 点以上なら合格、そうでなければ不合格
// カメのアルルがこの大学の入学試験を受けたところ、1 次試験が A 点、2 次試験が B 点でした。
// アルルがこの試験に合格したならば Yes、そうでなければ No を出力してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1次試験 A点, 2次試験 B点
	scanner.Scan()
	sliceStringOfScore := strings.Fields(scanner.Text())

	if len(sliceStringOfScore) != 2 {
		fmt.Println("入力形式が不正です")
		return
	}

	firstExamScore, err := strconv.Atoi(sliceStringOfScore[0])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	secondExamScore, err := strconv.Atoi(sliceStringOfScore[1])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	// 合格基準
	scanner.Scan()
	sliceStringOfPassScore := strings.Fields(scanner.Text())

	if len(sliceStringOfPassScore) != 2 {
		fmt.Println("入力形式が不正です")
		return
	}

	firstExamPassScore, err := strconv.Atoi(sliceStringOfPassScore[0])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	totalPassScoreRequirement, err := strconv.Atoi(sliceStringOfPassScore[1])
	if err != nil {
		fmt.Println("スコアは整数で入力してください")
		return
	}

	// 1 次試験の点数が X 点未満であれば、不合格
	if firstExamScore < firstExamPassScore {
		fmt.Println("No")
		return
	}

	// 1 次試験と 2 次試験の合計点が Y 点以上なら合格、そうでなければ不合格
	if firstExamScore + secondExamScore >= totalPassScoreRequirement {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}