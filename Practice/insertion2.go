package main

import "fmt"

func main()  {
	fmt.Println("Insertion sort partII")
	nums := []int{5, 2, 3, 1}
	insertionSort(nums)
	fmt.Println(nums)
}
func insertionSort(nums []int)  {
	for i := 1; i < len(nums); i++{
		insertion(nums, nums[i], i - 1)
	}
}

func insertion(nums []int, current int, previousOne int)  {
	for previousOne >= 0 && current < nums[previousOne] {
		nums[previousOne + 1] = nums[previousOne]
		previousOne--
	}

	nums[previousOne + 1] = current
}