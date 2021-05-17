package main

import "fmt"

func main()  {
	quickNum := []int{26, 5, 37, 1, 61, 11, 59, 15, 48, 19}
	fmt.Println("Before quick sort: ", quickNum)
	quickSort(quickNum, 0, len(quickNum) - 1)
	fmt.Println(quickNum)

}

func quickSort(num []int, left int ,right int)  {
	// 只剩一個元素時就回傳
	if left >= right {
		return
	}
	// 標記最左側元素作為 C.K.
	pivot := num[left]
	// 兩個游標分別從兩端相向移動，尋找合適的"支點"
	i := left
	j := right
	for i != j {
		// 右邊的游標向左移動，直到找到比 C.K. 的元素值小的
		for num[j] >= pivot && i < j {
			fmt.Println("num[j]", num[j])
			fmt.Println("j", j)
			j--
		}
		// 左側游標向右移動，直到找到比 C.K. 元素值大的
		for num[i] <= pivot && i < j {
			fmt.Println("num[i]", num[i])
			fmt.Println("i", i)
			i++
		}
		// 如果找到的兩個游標位置不統一，就游標位置元素的值，並繼續下一輪尋找
		// 此時交換的左右位置的值，右側一定不大於左側。可能相等但也會交換位置，所以才叫不穩定的排序算法
		if i < j {
			num[i], num[j] = num[j], num[i]
			fmt.Println(num)
		}
	}
	// 這時的 i 位置已經是我們要找的支點了，交換位置
	fmt.Println("=======")
	num[left], num[i] = num[i], num[left]
	// 按支點位置把原數列分成两段，再各自逐步縮小範圍排序
	quickSort(num, left, i - 1)
	quickSort(num, i + 1, right)
}

//func slice(num []int, left, right int) int {
//	high := num[right]
//	for i := left; i < right; i++ {
//		if num[i] < high {
//			num[i], num[left] = num[left], num[i]
//			left++
//		}
//	}
//	num[left], num[right] = num[right], num[left]
//
//	return left
//}
