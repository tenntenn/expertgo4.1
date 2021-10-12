// リスト 4.1.26 norm.NFC 定数と norm.NFD 定数の使用例
package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func main() {
	s := "é"

	// "é" "\u00e9"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 正準等価性に基づいて分解
	s = norm.NFD.String(s)

	// "é" "e\u0301"
	fmt.Printf("%[1]q %+[1]q\n", s)

	// 正準等価性に基づいて合成
	s = norm.NFC.String(s)

	// "é" "\u00e9"
	fmt.Printf("%[1]q %+[1]q\n", s)
}
