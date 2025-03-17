package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 2 つの正の整数 A,B が空白区切りで入力されます。 A と B のうち一の位が小さい方の値を出力してください。
// ただし、A と B の一の位は異なることが保証されています。
// 補足
// 整数の一の位はその数を 10 で割った余りと等しいです。
// たとえば 17 の一の位は 7 ですが、これは 17 を 10 で割った余りと一致します。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())
	if len(textSlice) != 2 {
		fmt.Println("2つの整数を入力してください")
		return
	}

	num1, err := strconv.Atoi(textSlice[0])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	num2, err := strconv.Atoi(textSlice[1])
	if err != nil {
		fmt.Println("入力は整数である必要があります。")
		return
	}

	num1LastDigit := num1 % 10
	num2LastDigit := num2 % 10

	if num1LastDigit < num2LastDigit {
		fmt.Println(num1)
	} else {
		fmt.Println(num2)
	}
}