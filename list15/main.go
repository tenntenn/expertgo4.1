// リスト 4.1.15 構造体の初期化(フィールド名を指定しない場合と、指定する場合)
package main

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	// フィールド名が指定されていないのでコンパイルエラー
	_ = &encoding.Encoder{transform.Nop}

	// OK
	_ = &encoding.Encoder{Transformer: transform.Nop}
}
