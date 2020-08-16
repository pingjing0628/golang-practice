package main

import "fmt"

// 4.5 撰寫原址函式以消除 []string 相鄰的重複
func main()  {
	s := []string{"A", "A", "B", "B", "B", "A", "C"}
	fmt.Println(removeDuplicate(s))
}

func removeDuplicate(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}

	return s
}