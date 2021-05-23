package main

import (
	"fmt"
)

func main()  {
	//A := []int{2, 2, 2, 2, 1, 2, -1, 2, 1, 3}
	A := []int{1, 2, 5, 19, -1, -2, -2, 5, 6, 7, 10, 18, 20, -2, -5}
	//A := []int{30 , 20 , 10}
	index := ascending(A)

	fmt.Println(index)
}

func ascending(A []int) int {
	i := 0 // Current index
	size := 0 // Calc size
	head := 0 // The start of begin ascending
	maxRecord := [3]int{} // Record the Biggest ascending size

	for i < len(A) - 1 {
		// 若後比前大
		if A[i] < A[i + 1] {
			// 若啟始點尚為 0, replace by i
			if head == 0 {
				head = i
			}
			// 若起始點至現在 index 比 原有 size 大, replace value in maxRecord
			if i + 1 - head > size {
				size = i + 1 - head
				maxRecord[0] = head
				maxRecord[1] = i + 1
				maxRecord[2] = size
			}
			// Next run
			i++
		}else{
			// 歸零
			head = 0
			i++
		}
	}

	//for i := 0; i < len(A)-1; i++{
	//	for j := i+1; j < len(A); j++{
	//		if A[i] < A[j] && j-i+1 > size{
	//			fmt.Println("A[i]", A[i])
	//			fmt.Println("A[j]", A[j])
	//			begin = i
	//			fmt.Println("begin", begin)
	//			size = j-i+1
	//			fmt.Println("size", size)
	//		}
	//	}
	//}
	return maxRecord[0]
}

