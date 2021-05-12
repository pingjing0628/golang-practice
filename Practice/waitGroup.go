package main

import (
	"fmt"
	"sync"
)

func foo()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar()  {
	for i := 1; i <= 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main()  {
	// add 2 thread in wait group
	wg.Add(2)
	go foo()
	go bar()
	// 等待 WaitGroup 裡的 Thread 執行完畢再繼續進行
	wg.Wait()
}