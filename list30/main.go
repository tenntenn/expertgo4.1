// リスト 4.1.30 アルファベットを大文字に変換
package main

import (
	"fmt"

	"golang.org/x/text/runes"
)

func main() {
	t := runes.Map(func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	})
	// HELLO, WORLD
	fmt.Println(t.String("Hello, World"))
}
