// リスト 4.1.17 Shift_JIS の CSV ファイルを UTF-8 に変換して読み込む
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/japanese"
)

func main() {
	if err := printCSV("sample.csv"); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func printCSV(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Shift_JISとして読み込む
	dec := japanese.ShiftJIS.NewDecoder()
	cr := csv.NewReader(dec.Reader(f))
	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// UTF-8に変換されているので表示しても
		// 文字化けしない
		fmt.Println(rec)
	}

	return nil
}
