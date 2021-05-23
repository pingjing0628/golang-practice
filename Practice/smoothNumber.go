package main

// you can also use imports, for example:
// import "fmt"
import (
	"fmt"
	"sort"
)

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
func main()  {
	a := Solution(4)
	fmt.Println(a)
}

// 3-smooth
// smooth number 可以因數分解為小質數乘積的正整數
func Solution(N int) int {
	// write your code in Go 1.4
	Total := 16000000
	if N <= 200 && N > 1 {
		A := []int{}
		multiply2 := 1 // 預設為 1
		for multiply2 < Total {
			multiply2And3 := multiply2 // 預設為 multiply2
			for multiply2And3 < Total { // 執行 multiply 2 * 3
				A = append(A, multiply2And3) // 加入陣列 A
				multiply2And3 *= 3 // *3
			}
			multiply2 *= 2 // *2
		}
		sort.Ints(A) // 排序數值

		return A[N]
	} else {
		return 0
	}
}


