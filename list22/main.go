// リスト 4.1.22 文字コード変換と全角半角変換を行う
package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"golang.org/x/text/width"
)

func main() {
	if err := foldShiftJISFile("sample.txt"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

// Shift_JISのファイルの全角英数などは半角に、半角カナなどは全角にする
func foldShiftJISFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JISからUTF-8に変換してから
	// 全角英数などは半角に、半角カナなどは全角にする
	dec := japanese.ShiftJIS.NewDecoder()
	t := transform.Chain(dec, width.Fold)
	r := transform.NewReader(f, t)
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		return err
	}

	return nil
}
