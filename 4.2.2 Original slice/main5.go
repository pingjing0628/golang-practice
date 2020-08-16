package main

import (
	"fmt"
	"unicode/utf8"
)

// 4.7 Modify reverse function to reverse UTF8 char
func main()  {
	s := "Test For Reverse"

	fmt.Println(s)
	fmt.Println(string(reverseUTF8([]byte(s))))
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse2(b[i:i+size])
		i += size
	}

	reverse2(b)

	return b
}

func reverse2(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1{
		b[i], b[j] = b[j], b[i]
	}

	return b
}