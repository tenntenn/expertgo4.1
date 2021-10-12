// リスト 4.1.18 East Asian Width 特性を取得
package main

import (
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	// 全角の5、半角のア、全角のア、半角のA、ギリシア文字のアルファ
	rs := []rune{'5', 'ア', 'ア', 'A', 'α'}
	fmt.Println("rune\tWide\tNarrow\tFolded\tKind")
	fmt.Println("--------------------------------------------------")
	for _, r := range rs {
		p := width.LookupRune(r)
		w, n, f, k := p.Wide(), p.Narrow(), p.Folded(), p.Kind()
		fmt.Printf("%2c\t%2c\t%3c\t%3c\t%s\n", r, w, n, f, k)
	}
}
