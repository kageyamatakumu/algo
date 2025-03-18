package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// 問題文
// 夏季オリンピックは、原則として西暦が 4 の倍数の年に開催されます。
// ただし、2020 年に予定されていた東京オリンピックは延期され、2021 年に開催されました。
// この延期が 1948 年から 2023 年現在までで唯一の例外です。
// 整数 N(1948≤N≤2023) が与えられます。西暦 N 年にオリンピックが開催されたならば Yes、そうでないならば No を出力してください。

func main() {
	// 入力可能な最小年と最大年
	const validYearMin = 1948
	const validYearMax = 2023

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	year, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("整数で入力してください")
		return
	}

	if year < validYearMin || year > validYearMax {
		fmt.Println("オリンピック開催期間内で入力してください")
		return
	}

	// 特例の処理
	if year == 2020 {
		fmt.Println("No")
		return
	} else if year == 2021 {
		fmt.Println("Yes")
		return
	}

	// 4の倍数であるか判定
	if year % 4 == 0  {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}