// リスト 4.1.25 文字列“Cafe\u0301”(Café)を書記素クラスタに分解
package main

import (
	"fmt"

	"github.com/rivo/uniseg"
)

func main() {
	gr := uniseg.NewGraphemes("Cafe\u0301")
	// C [43] a [61] f [66] é [65 301]
	for gr.Next() {
		fmt.Printf("%s %x ", gr.Str(), gr.Runes())
	}
}
