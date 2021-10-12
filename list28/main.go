// リスト 4.1.28 文字列中のカタカナをすべて全角に変換
package main

import (
	"fmt"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/width"
)

func main() {
	// カタカナであれば全角にする
	t := runes.If(runes.In(unicode.Katakana), width.Widen, nil)
	// 5アアAα
	fmt.Println(t.String("5アアAα"))
}
