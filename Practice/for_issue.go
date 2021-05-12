package main

import "fmt"

type student struct {
	Name string
	Age int
}

func parseStudent()  {
	m := make(map[string]*student)
	studs := []student{
		{Name: "wang", Age: 24},
		{Name: "hui", Age:25},
		{Name: "yu", Age: 26},
	}

	fmt.Println(m)

	// Issue
	for _, stu := range studs {
		m[stu.Name] = &stu // 實際上一致指向同一個指標，最終該指標的值為遍歷的最後一個 struct 的值拷貝
	}

	for k, v := range m{
		println(k, "=>", v.Name)
	}

	// Correct
	for i := 0; i < len(studs); i++ {
		m[studs[i].Name] = &studs[i]
	}

	for k, v := range m{
		println(k, "=>", v.Name)
	}
	fmt.Println(m)
}

func main()  {
	parseStudent()
}
