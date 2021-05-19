package main

import (
	"fmt"
)

func main()  {
	fmt.Println("Quicksort partII")
	quickNums := []int{5, 2, 3, 1}
	quicksort(quickNums, 0, len(quickNums) - 1)
	fmt.Println(quickNums)
}

func quicksort(quickNums []int, low int, high int)  {
	if low > high {
		return
	}

	// Tag C.K
	pivot := quickNums[low]
	i := low
	j := high

	for i != j {
		for quickNums[j] >= pivot && i < j {
			j--
		}
		for quickNums[i] <= pivot && i < j {
			i++
		}
		if i < j {
			quickNums[i], quickNums[j] = quickNums[j], quickNums[i]
		}
	}

	quickNums[low], quickNums[i] = quickNums[i], quickNums[low]
	quicksort(quickNums, low, i - 1)
	quicksort(quickNums, i + 1, high)

}
