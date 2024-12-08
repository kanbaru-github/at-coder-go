package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	// RLE (Run Length Encoding) の実装
	type pair struct {
		char  rune
		count int
	}
	var rle []pair

	for _, c := range s {
		if len(rle) > 0 && rle[len(rle)-1].char == c {
			rle[len(rle)-1].count++
		} else {
			rle = append(rle, pair{char: c, count: 1})
		}
	}

	// '1' の連続部分の k 番目で swap 処理を行う
	oneCount := 0
	for i := 0; i < len(rle); i++ {
		if rle[i].char == '1' {
			oneCount++
			if oneCount == k && i > 0 {
				// 前後をスワップ
				rle[i-1], rle[i] = rle[i], rle[i-1]
			}
		}
	}

	// 答えを構築
	var ans strings.Builder
	for _, p := range rle {
		ans.WriteString(strings.Repeat(string(p.char), p.count))
	}
	fmt.Println(ans.String())
}
