// リスト 4.1.19 通常半角で表す文字は半角に、通常全角で表す文字は全角に変換
package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	// 全角の5、半角のア、全角のア、半角のA、ギリシア文字のアルファ
	for _, r := range width.Fold.String("5アアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}
