package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// カメのアルルは A 時に出社し、B 時に退社しました。
// その途中、C 時から D 時までの間休憩を取っていました。
// A,B,C,D の値が標準入力から与えられます。
// アルルが働いていた時間を計算し出力するプログラムを作成してください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 4 {
		fmt.Println("入力は4つの整数である必要があります。")
		return
	}

	A, err := strconv.Atoi(textSlice[0])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	B, err := strconv.Atoi(textSlice[1])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	C, err := strconv.Atoi(textSlice[2])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	D, err := strconv.Atoi(textSlice[3])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	if A < 0 || A > 23 || B < 0 || B > 23 || C < 0 || C > 23 || D < 0 || D > 23 {
		fmt.Println("入力は0以上23以下の整数である必要があります。")
		return
	}

	if A > B {
		fmt.Println("出社時間は退社時間より前でなければなりません。")
		return
	}

	if C > D {
		fmt.Println("休憩時間の開始時間は終了時間より前でなければなりません。")
		return
	}

	restTime := D - C


	fmt.Println(B - A - restTime)
}