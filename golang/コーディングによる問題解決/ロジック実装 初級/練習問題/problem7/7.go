package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 文字列 S が 1 行目に、2 つの正の整数 N,M が空白区切りで 2 行目に入力されます。 S の前から N 番目の文字と、前から M 番目の文字を入れ替えた文字列を出力してください。
// ただしここでは、文字列 S の先頭の文字は 1 文字目であるとします。

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()
	length := len(text)

	scanner.Scan()
	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 2 {
		fmt.Println("入力は2つの整数である必要があります。")
		return
	}

	N, err := strconv.Atoi(textSlice[0])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	M, err := strconv.Atoi(textSlice[1])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	if N <= 0 || M <= 0 || N > length || M > length {
		fmt.Println("入力した数字が文字列の範囲外です。")
		return
	}

	runes := []rune(text)
	runes[M-1], runes[N-1] = runes[N-1], runes[M-1]
	fmt.Println(string(runes))
}