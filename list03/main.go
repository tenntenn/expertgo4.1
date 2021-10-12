// リスト 4.1.3 複数パターンの置換
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 郷 → Go、入れば → 入っては
	r := strings.NewReplacer("郷", "Go", "入れば", "入っては")
	s := r.Replace("郷に入れば郷に従え")
	// Goに入ってはGoに従え
	fmt.Println(s)
}
