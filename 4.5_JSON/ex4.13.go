package main

import (
	"fmt"
	"github.com/pingjing0628/golang-practice/4.5_JSON/omdb"
	"os"
)

func main()  {
	// Get poster according to the input value
	err := omdb.GetPoster(os.Stdout, os.Args[1:])

	if err != nil {
		fmt.Fprintf(os.Stderr, "ex13: %v\n", err)
		os.Exit(1)
	}
}
