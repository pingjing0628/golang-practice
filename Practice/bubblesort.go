package main

import "fmt"

func main()  {
	bubbleNum := []int{23, 5, 6, 77, 13}
	fmt.Println("Before bubble sort: ", bubbleNum)
	bubbleSort(bubbleNum)
}

func bubbleSort(num []int)  {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(num) - 1; i++{ // 長度-1 因為從0開始
			if num[i] > num[i + 1]{
				num[i+1], num[i] = num[i], num[i+1]
				swapped = true
			}
		}
	}
	fmt.Println(num)
}