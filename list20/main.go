// リスト 4.1.20 全角文字を半角文字に変換
package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	for _, r := range width.Narrow.String("5アアAα") {
		p := width.LookupRune(r)
		fmt.Printf("%c: %s\n", r, p.Kind())
	}
}
