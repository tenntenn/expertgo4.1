// リスト 4.1.2 文字列の置換(第 2 引数で指定した文字列すべてを置換する)
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Replace("郷に入っては郷に従え", "郷", "Go", -1)
	// Goに入ってはGoに従え
	fmt.Println(s)
}
