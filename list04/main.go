// リスト 4.1.4 小文字の英字を大文字の英字に変換
package main

import (
	"fmt"
	"strings"
)

func main() {
	toUpper := func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	}
	s := strings.Map(toUpper, "Hello, World")
	// HELLO, WORLD
	fmt.Println(s)
}
