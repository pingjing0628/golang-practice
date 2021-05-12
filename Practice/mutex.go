package main

import (
	"fmt"
	"sync"
	"time"
)

var money int = 1500
var wgg sync.WaitGroup
var mu sync.Mutex

func withdraw()  {
	mu.Lock()
	balance := money
	time.Sleep(3000 * time.Millisecond)
	balance -= 1000
	money = balance
	mu.Unlock()
	fmt.Println("After withdrawing $1000, balance: ", money)
	wgg.Done()
}

func main()  {
	fmt.Println("We have $1500")
	wgg.Add(2)
	go withdraw()
	go withdraw()
	wgg.Wait()
}
