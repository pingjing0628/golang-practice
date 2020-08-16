package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 4.6 撰寫原址函式壓縮 UTF8 編碼的 []byte 相鄰的 Unicode 空白
func main()  {
	s := "1   +   1  = 2"

	fmt.Println(s)
	fmt.Println(string(removeDuplicateSpace([]byte(s))))
}

func removeDuplicateSpace(b []byte) []byte {
	var buf bytes.Buffer

	for i := 0; i < len(b); {
		r, size := utf8.DecodeRuneInString(string(b[i:]))

		if unicode.IsSpace(r) {
			nextRune, _ := utf8.DecodeLastRuneInString(string(b[i+size:]))

			if ! unicode.IsSpace(nextRune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}

		i += size
	}

	return buf.Bytes()
}