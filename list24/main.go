// リスト 4.1.24 UTF-8 を Unicode コードポイントにデコード
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("世界")
	// '世':3 '界':3
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%q:%v ", r, size)
		b = b[size:]
	}
}
