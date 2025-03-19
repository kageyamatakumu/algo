package main

import (
	"bufio"
	"fmt"
	"os"
)

// 問題文
// カメのアルルは卓球を 5 試合行いました。
// 試合結果は o と x のみからなる長さ 5 の文字列 S で表され、S の各文字は各試合の勝敗を表します。
// たとえば S が oxxox の場合、アルルの試合結果が順に 勝ち・負け・負け・勝ち・負け であったことを意味します。
// アルルの勝利数を求めてください。

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	var victory int

	for _, v := range input {
		if v == 'o' {
			victory++
		}
	}

	fmt.Println(victory)
}