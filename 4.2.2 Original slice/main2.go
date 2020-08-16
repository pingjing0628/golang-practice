package main

import "fmt"

// 4.4 rotate
func main()  {
	s := []int{0, 1, 2, 3, 4, 5}

	// rotate s left by two postions
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, n int) {
	num := n % len(s)

	double := append(s, s[:num]...)

	copy(s, double[num:num+len(s)])
}
