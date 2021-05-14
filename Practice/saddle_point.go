package main

import "fmt"

// A[i,j]是第i行中值最小的元素,且又是第j列中值最大的元素，則稱之為該矩陣的一個馬鞍點
func main()  {
	fmt.Println("Saddle point in two-dimension array")
	m := 3
	n := 3
	A := [3][3]int{{1, 7, 3}, {5, 4, 6}, {17, 18, 9}}
	fmt.Println(A)
	saddle(m, n, A)
}
func saddle(m int, n int, A [3][3]int)  {
	var hMin[10]int
	var lMax[10]int

	// 先找出每行中最小的
	for i := 0; i < m; i++ {
		// 預設最小值為每列的頭
		hMin[i] = A[i][0]
		for j := 0; j < n; j++{
			if A[i][j] < hMin[i]{
				hMin[i] = A[i][j]
			}
		}

	}

	// 在找出每列中最大的
	for j := 0; j < n; j++{
		// 預設最大的是每行的頭
		lMax[j] = A[0][j]
		for i := 0; i < m; i++ {
			if A[i][j] > lMax[j]{
				lMax[j] = A[i][j]
			}
		}

	}

	for i := 0; i < m; i++{
		if hMin[i] == lMax[i]{
			fmt.Println(hMin[i])
		}
	}

}
