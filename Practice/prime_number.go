package main

import "fmt"

func main()  {
	fmt.Println("print prime number smaller than 100")
	for i := 2; i <= 100; i++ {
		bool := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				bool = false
				break
			}
		}
		if bool {
			fmt.Println(i, " ")
		}
	}
}
