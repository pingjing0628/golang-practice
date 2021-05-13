package main

import (
	"fmt"
)

func main()  {
	selectionNum := []int{23, 5, 6, 77, 13}
	fmt.Println("Before selection sort: ", selectionNum)
	selectionSort(selectionNum)
}

func selectionSort(num []int)  {
	for i := 0; i < len(num) - 1; i++ { //不做最後回合
		m := i // m 記錄最小值位置
		for j := i + 1; j < len(num); j++{
			if num[j] < num[m] {
				m = j // j 從 i+1...n 找 min, 並將"地址"丟給 m
			}
		}
		if i != m {  // means i 是最小值
			num[i], num[m] = num[m], num[i]
		}
	}
	fmt.Println("Sort over", num)
}
