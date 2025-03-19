package main

import "fmt"

// 問題文
// 次の処理を i=1,2,…,100 について行ってください。
// i 行目には整数 2i を出力する。

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i * 2)
	}
}