package main

import "fmt"

func main() {
	// 1. Slice 會動態改變其長度
	// a[1:] or a[:] means 指向整個陣列
	// Example: months[1:13], [4:7] 4 means from 4 to 7 個位置之值, len 為 3 , cap 為 9
	a := [...]int{0, 1, 2, 3, 4, 5}

	// 翻轉 int  slice
	reverse(a[:]) // 全部
	fmt.Println(a)

	b := []int{0, 1, 2, 3, 4, 5}

	// 向左旋轉兩個位置

	reverse(b[:2])
	fmt.Println(b)

	reverse(b[2:])
	fmt.Println(b)

	reverse(b)
	fmt.Println(b)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 2. Slice is not like Array, it cannot be compared.
// Which means it cannot test two slice has the same element by using == .
// The only valid comparsion in slice is compare with the nil.
// If you want to test a slice is nil or not, you should use len(s) == 0 not s == nil.

// 3. make function 建構指定 型別、長度、容量的 slice. 容量可省略，則容量 = 長度
// make([]T, len)
// make([]T, len, cap) = make([]T,
