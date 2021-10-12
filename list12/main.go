// リスト 4.1.12 transform.NewWriter 関数の使用例
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/transform"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	// 変数rはio.Readerインタフェースを実装した型
	r := strings.NewReader("Hello, World")

	// transform.Nop変数は何も変換を行わないtransform.Transformer
	tw := transform.NewWriter(os.Stdout, transform.Nop)

	// 変数twは*transform.Writer型
	_, err := io.Copy(tw, r)
	if err != nil {
		return err
	}

	return nil
}
