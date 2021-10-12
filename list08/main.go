// リスト 4.1.8 正規表現を使った置換(キャプチャした文字列を展開しない)
package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		return err
	}

	s := re.ReplaceAllLiteralString("1986年01月12日", "${2}/${3} ${1}")
	// ${2}/${3} ${1}
	fmt.Println(s)

	return nil
}
