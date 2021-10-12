// リスト 4.1.6 []byte 型の値の置換(すべての置換対象を置換する)
package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 0x0B → 0xFF
	b := bytes.ReplaceAll([]byte{0x0A, 0x0B, 0x0C}, []byte{0x0B}, []byte{0xFF})
	// 0A FF 0C
	fmt.Printf("% X\n", b)
}
