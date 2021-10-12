// リスト 4.1.9 正規表現にマッチした文字列を、関数で指定したルールで変換
package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	re, err := regexp.Compile(`(^|_)[a-zA-Z]`)
	if err != nil {
		return err
	}

	s := re.ReplaceAllStringFunc("hello_world", func(s string) string {
		return strings.ToUpper(strings.TrimLeft(s, "_"))
	})
	// HelloWorld
	fmt.Println(s)

	return nil
}
