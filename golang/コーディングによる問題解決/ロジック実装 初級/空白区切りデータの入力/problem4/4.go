package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 問題文
// 3 つの整数 A,B,C が空白区切りで入力されます。3 つの整数の平均値を整数で出力してください。
// ただし、答えは整数になることが保証されています。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	textSlice := strings.Fields(scanner.Text())

	if len(textSlice) != 3 {
		fmt.Println("入力は3つの整数である必要があります。")
		return
	}

	var sum int

	for _, numStr := range textSlice {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("入力は整数である必要があります。")
			return
		}
		sum += num
	}

	fmt.Println(sum / 3)
}