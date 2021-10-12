// リスト 4.1.1 文字列の置換
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Replace("郷に入っては郷に従え", "郷", "Go", 1)
	// Goに入っては郷に従え
	fmt.Println(s)
}
