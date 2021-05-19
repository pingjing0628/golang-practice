package main

import (
	"fmt"
)

func main()  {
	//A := []int{2, 2, 2, 2, 1, 2, -1, 2, 1, 3}
	A := []int{1, 2, 5, 19, -1, -2, -2, 5, 6, 7, 10, 18, 20, -2, -5}
	index := ascending(A)

	fmt.Println(index)
}

func ascending(A []int) int {
	begin := 0
	size := 0
	fmt.Println(len(A))
	for i := 0; i < len(A)-1; i++{
		for j := i+1; j < len(A); j++{
			if A[i] < A[j] && j-i+1 > size{
				fmt.Println("A[i]", A[i])
				fmt.Println("A[j]", A[j])
				begin = i
				fmt.Println("begin", begin)
				size = j-i+1
				fmt.Println("size", size)
			}
		}
	}
	return begin
}

