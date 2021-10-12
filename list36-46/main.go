// 任意のバイト列の変換
package main

import (
	"bytes"
	"io"
	"os"
	"strings"

	"golang.org/x/text/transform"
)

// リスト 4.1.36 Replacer 型の使用例(任意のバイト列変換)
func main() {
	// "郷"を"Go"に変換する
	var t *Replacer = NewReplacer([]byte("郷"), []byte("Go"))
	w := transform.NewWriter(os.Stdout, t)
	// Goに入ってはGoに従え
	io.Copy(w, strings.NewReader("郷に入っては郷に従え"))
}

// リスト 4.1.37 Replacer 型の構造体
type Replacer struct {
	old, new []byte
	// 前回書き込めなかった分（リスト 4.1.40）
	preDst []byte
	// 前回余ったold分（リスト 4.1.43）
	preSrc []byte
}

func NewReplacer(old, new []byte) *Replacer {
	return &Replacer{
		old: old,
		new: new,
	}
}

func (r *Replacer) Transform(dst, src []byte, atEOF bool) (int, int, error) {

	// リスト 4.1.44 preSrc フィールドに関する前処理と後処理

	_src := src
	if len(r.preSrc) > 0 {
		_src = make([]byte, len(r.preSrc)+len(src))
		copy(_src, r.preSrc)
		copy(_src[len(r.preSrc):], src)
	}

	nDst, nSrc, preSrc, err := r.transform(dst, _src, atEOF)

	if nSrc < len(r.preSrc) {
		r.preSrc = r.preSrc[nSrc:]
		nSrc = 0
	} else {
		nSrc -= len(r.preSrc)
		r.preSrc = preSrc
	}

	return nDst, nSrc, err
}

func (r *Replacer) transform(dst, src []byte, atEOF bool) (nDst, nSrc int, preSrc []byte, err error) {
	// リスト 4.1.41
	// 前回書き込めなかった分を書き込む
	if len(r.preDst) > 0 {
		n := copy(dst, r.preDst)
		nDst += n
		r.preDst = r.preDst[n:]
		// それでもまだ足りない場合
		if len(r.preDst) > 0 {
			err = transform.ErrShortDst
			return
		}
	}

	// リスト 4.1.38 old フィールドの長さが 0 の場合の処理
	// r.oldが空であれば、そのままコピー
	if len(r.old) == 0 {
		n := copy(dst[nDst:], src)
		nDst += n
		nSrc += n
		return
	}

	// リスト 4.1.39 変換対象のバイト列が見つかった場合の処理
	for {
		i := bytes.Index(src[nSrc:], r.old)

		// リスト 4.1.42 変換対象のバイト列が見つからなかった場合の処理
		// 見つからなかった場合
		if i == -1 {
			n := len(src[nSrc:])

			// リスト 4.1.45
			// srcの末尾がr.oldの前方部分で終わってる場合
			var w int
			if !atEOF { // まだ次で読み込める余地ある
				// srcの末尾とr.oldが同じ分の長さ
				w = overlapWidth(src[nSrc:], r.old)
				if w > 0 {
					// 足りなかった分はコピーする分から一旦除外しておく
					n -= w
					err = transform.ErrShortSrc
				}
			}

			m := copy(dst[nDst:], src[nSrc:nSrc+n])
			nDst += m
			nSrc += m
			if m < n {
				err = transform.ErrShortDst
				return
			}

			// リスト 4.1.45
			// 次のsrcの先頭に追加しておく
			preSrc = r.old[:w]
			// 読み込んだことにする
			nSrc += w

			return
		}

		// 見つけたところまでをコピーして書き込む
		n := copy(dst[nDst:], src[nSrc:nSrc+i])
		nDst += n
		nSrc += n
		if n < i {
			err = transform.ErrShortDst
			return
		}

		// 置換する文字をコピーして書き込む
		n = copy(dst[nDst:], r.new)
		nDst += n
		nSrc += len(r.old)

		// リスト 4.1.41
		// r.newが長くてdstが足りない場合は次回に持ち越し
		if n < len(r.new) {
			r.preDst = r.new[n:]
			err = transform.ErrShortDst
			return
		}
	}
}

// リスト 4.1.46 overlapWidth 関数
// aとbで先頭からマッチする長さ
// 例: a:[1, 2, 3], b:[1, 2] -> 2
func overlapWidth(a, b []byte) int {

	// aとbで短い方の長さ
	w := len(a)
	if w > len(b) {
		w = len(b)
	}

	// 短くしながらマッチするまで
	for ; w > 0; w-- {
		if bytes.Equal(a[len(a)-w:], b[:w]) {
			return w
		}
	}

	return 0
}

func (r *Replacer) Reset() {
	r.preDst = nil // （リスト 4.1.40）
	r.preSrc = nil // （リスト 4.1.43）
}
