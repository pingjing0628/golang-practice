package main

import "fmt"

func main()  {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y

		// 輸出的每個容量改變表示一個分配與複製
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	if zlen <= cap(x) {
		// 有空間供成長，擴張 slice
		z = x[:zlen]
	} else {
		// 空間不夠，分配新陣列
		// 倍增，平攤線性複雜度
		zcap := zlen

		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)
		copy(z, x)
	}

	z[len(x)] = y
	return z
}
