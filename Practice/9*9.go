package main

import (
	"fmt"
)

func main() {
	fmt.Println("9*9")
	for i := 1; i <=9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Println(i, "*", j , "=", i*j)
		}
	}
}