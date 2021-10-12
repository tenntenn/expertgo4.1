// アルファベットの小文字から大文字への変換
package main

import (
	"io"
	"os"
	"strings"

	"golang.org/x/text/transform"
)

// リスト 4.1.32 transform.NopResetter 型を埋め込む
type Upper struct{ transform.NopResetter }

func (Upper) Transform(dst, src []byte, atEOF bool) (
	nDst, nSrc int, err error) {

	// リスト 4.1.33 引数 dst の長さが不十分な場合のエラー処理
	// 末尾ではないのにdstが足りない場合はErrShortDstを返す
	if len(dst) == 0 && !atEOF {
		err = transform.ErrShortDst
		return
	}

	// リスト 4.1.34 バイト列の変換処理
	for {
		// srcをすべて処理し終えた、またはdstが足りなくなった場合
		if len(src) <= nSrc || len(dst) <= nDst {
			return
		}

		// 小文字から大文字への変換
		if 'a' <= src[nSrc] && src[nSrc] <= 'z' {
			dst[nDst] = src[nSrc] - 'a' + 'A'
		} else {
			dst[nDst] = src[nSrc]
		}

		// 処理したバイト数分だけ進める
		nSrc++
		nDst++
	}

}

// リスト 4.1.35 Upper 型の使用例(アルファベットの大文字変換)
func main() {
	var upper Upper
	sr := strings.NewReader("Hello, World")
	r := transform.NewReader(sr, &upper)

	// HELLO, WORLD
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		panic(err)
	}
}
