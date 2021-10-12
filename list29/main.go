// リスト 4.1.29 アクサンテギュなどを削除
package main

import (
	"fmt"
	"os"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	removeMn := runes.Remove(runes.In(unicode.Mn))
	t := transform.Chain(norm.NFD, removeMn, norm.NFC)

	s, _, err := transform.String(t, "résumé")
	if err != nil {
		return err
	}

	// resume
	fmt.Println(s)

	return nil
}
