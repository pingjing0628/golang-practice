package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ex4.9 To count character times in import file
func main()  {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Must have 1 argument.")
		os.Exit(1)
	}

	counts := make(map[string]int) // counts of words

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(file)

	input.Split(bufio.ScanWords) // 將輸入拆分成字
	for input.Scan() {
		word := strings.ToLower(input.Text())
		counts[word]++
	}

	fmt.Printf("word\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}

