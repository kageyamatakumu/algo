package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// N 個の正の整数 A0,A1,…,AN-1が与えられます。N 個の整数の合計値を求めてください。
// また、入力される値は次の制約を満たします。
// N は 1 以上 1000 以下の整数
// Ai​は 1 以上 1000 以下の整数 (0≤i≤N−1)

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

	// Aの配列を受け取る
	scanner.Scan()
	inputA := scanner.Text()
	arrAStr := strings.Fields(inputA)

	if len(arrAStr) != N {
		fmt.Println("Nで入力した個数を入力してください")
		return
	}

	var answer int

	for _, v := range arrAStr {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("1 以上 1000 以下の整数で入力してください")
			return
		}

		if num < 1 || num > 1000 {
			fmt.Println("1 以上 1000 以下の整数で入力してください")
			return
		}

		answer += num
	}

	fmt.Println(answer)
}