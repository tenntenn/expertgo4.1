// リスト 4.1.23 Unicode コードポイントを UTF-8 にエンコード
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	// [228 184 150] "世" 3
	fmt.Printf("%v %q %d", buf, string(buf), n)
}
