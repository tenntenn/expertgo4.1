// リスト 4.1.5 小文字を大文字に変換(unicode.ToUpper 関数)
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := strings.Map(unicode.ToUpper, "Hello, World")
	// HELLO, WORLD
	fmt.Println(s)
}
