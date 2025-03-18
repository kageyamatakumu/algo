package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 問題文
// 正の整数 N が与えられます。
// N が 3 の倍数ならば、Fizz と出力してください
// N が 5 の倍数ならば、Buzz と出力してください
// ただし、N が 3 の倍数でも 5 の倍数でもあるならば、FizzBuzz と出力してください
// それ以外の場合は、そのままの整数 N を出力してください

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("数字で入力してください")
		return
	}

	if num <= 0 {
		fmt.Println("正の整数で入力してください")
		return
	}

	if num % 15 == 0 {
		fmt.Println("FizzBuzz")
	} else if num % 3 == 0 {
		fmt.Println("Fizz")
	} else if num % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}

}