// リスト 4.1.27 norm.NFKC 定数と norm.NFKD 定数の使用例
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	s := "ゴ"

	// "ゴ" "\u30b4"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 互換等価性に基づいて分解
	s = norm.NFKD.String(s)
	// "コ ゙" "\u30b3¥u3099"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 互換等価性に基づいて合成
	s = norm.NFKC.String(s)
	// "ゴ" "\u30b4"
	fmt.Printf("%[1]q %+[1]q\n", s)
}
