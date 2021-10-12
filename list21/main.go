// リスト 4.1.21 半角文字を全角文字に変換
package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	for _, r := range width.Widen.String("5アアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}
