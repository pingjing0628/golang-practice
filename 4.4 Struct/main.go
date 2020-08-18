package main

import (
	"fmt"
	"time"
)

// struct 中的每個值稱為欄位
// struct 型別通常是出現在像是 Employee 這樣的具名型別宣告中
type Employee struct {
	ID 			  int
	Name, Address string
	DoB 		  time.Time
	Position 	  string
	Salary 		  int
	ManagerID 	  int
}

var dilbert Employee

func main()  {
	dilbert.Salary -= 50000 // 不想加班，降薪

	// Use pointer
	position := &dilbert.Position
	*position = "Senior " + *position // 拍馬屁，加薪

	// 使用函式回傳 Employee 指標
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 // Fired
}

//
func EmployeeByID(id int) *Employee {
	/* ... */
	//return
}