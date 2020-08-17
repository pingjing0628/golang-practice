package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Count Unicode char times
// Ex4.8 use unicode.IsLetter to count char, number
func main()  {
	counts := make(map[string]int) // count Unicode char times
	//var utflen [utf8.UTFMax + 1]int // count UTF-8 encode length
	invalid := 0 // count invalid UTF-8 char

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // return rune, nbytes, error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(r) {
			counts["Letter"]++
		}

		if unicode.IsNumber(r) {
			counts["Number"]++
		}
		//counts[r]++
		//utflen[n]++
	}

	//fmt.Printf("rune\tcount\n")

	fmt.Printf("kind\tcount\n")

	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}

	//fmt.Print("\nlen\tcount\n")

	//for i, n := range utflen {
	//	if i > 0 {
	//		fmt.Printf("%d\t%d\n", i, n)
	//	}
	//}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
