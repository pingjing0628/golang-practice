package main

import "fmt"

func insertSort(num []int) {
	for i := 1; i < len(num); i++ {
		insert(num[i], num, i - 1)
	}
}

func insert(current int, num []int, previousOne int)  {
	for previousOne >= 0 && current < num[previousOne] {
		num[previousOne + 1] = num[previousOne]
		previousOne--
	}
	num[previousOne + 1] = current
}

func main()  {
	num := []int{26, 5, 33, 17, 2}
	fmt.Println("Before sort: ", num)
	insertSort(num)
	fmt.Println(num)
}