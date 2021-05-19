package main

// you can also use imports, for example:
// import "fmt"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	Total := 200
	A := []int{}
	power2 := 1
	for power2 < Total {
		power2And3 := power2
		for power2And3 < Total {
			A = append(A, power2And3)
			power2And3 *= 3
		}
		power2 *= 2
	}
	sort.Ints(A)

	return A[N]
}

