package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main()  {
	c := []byte{}
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		c = []byte(input.Text())
	}

	var shaCase string

	if len(os.Args[1:]) == 1 {
		shaCase = os.Args[1]
	}

	switch shaCase {
		case "SHA384":
			fmt.Printf("%x\n", sha512.Sum384(c))
		case "SHA512":
			fmt.Printf("%x\n", sha512.Sum512(c))
		default:
			fmt.Printf("%x\n", sha256.Sum256(c))



	}
}