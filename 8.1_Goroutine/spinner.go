package main

import (
	"fmt"
	"time"
)

func main()  {
	// 設定等待時間
	go spinner(100 * time.Millisecond)

	const n = 45
	// main goroutine calculation
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// 建立 loading 圖示
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x - 1) + fib(x - 2)
}