package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func main()  {
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	i := 1
	i++
	j := i
	fmt.Println(j)

	var str = "hello 你好"
	// 5 word 1 blank 2 chinese word = 8
	// golang string 底層是透過 byte 陣列實現，中文字元 unicode 下佔2個位元組，在 utf-8 下佔3個位元組，而 golang 預設編碼正好是 utf-8
	fmt.Println("len(str):", len(str))

	// Sol1
	// golang中的unicode/utf8包提供了用utf-8獲取長度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	// Sol2
	// 通過rune型別處理unicode字元
	fmt.Println("rune:", len([]rune(str)))

	// byte 等同於int8，常用來處理ascii字元
	// rune 等同於int32,常用來處理unicode或utf-8字元
}
